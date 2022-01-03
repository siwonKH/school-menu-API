package main

import (
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"log"
	"net/url"
	"schoolMenuApi/model/apiResponse"
	"schoolMenuApi/request"
	"strings"
	"time"
)

type RecentSchool struct {
	Num        string
	SchoolName string
	AptCode    string
	SchoolCode string
	Date       string
	MenuData   apiResponse.Menu
}

// Cache in RAM!! wow
var recentSchoolsSize = 100
var recentSchools [101]RecentSchool
var recentSchoolsCnt = -1
var recentSchoolsTomorrow [101]RecentSchool

func prepareTomorrow() {
	log.Println("Starting cache renew and save")
	tomorrow := time.Now().Add(time.Hour * 23).Format("20060102")
	for i := 0; i < recentSchoolsSize; i++ {
		start := time.Now()
		MenuData := request.SearchMenu(recentSchools[i].AptCode, recentSchools[i].SchoolCode, tomorrow)
		recentSchoolsTomorrow[i] = recentSchools[i]
		recentSchoolsTomorrow[i].MenuData = MenuData
		recentSchoolsTomorrow[i].Date = tomorrow
		elapsed := time.Since(start)
		log.Printf("Saved %s in %s", recentSchools[i].SchoolName, elapsed)
		time.Sleep(1 * time.Second)
	}
	log.Println("Done")
}

func setTomorrowCache() {
	time.Sleep(59*time.Second + 500)
	recentSchools = recentSchoolsTomorrow
	log.Println("copied 'recentSchoolsTomorrow' to 'recentSchools'")
}

func apiMainProcess(schoolName string, decodedSchoolName string, dateStr string, num string, c *fiber.Ctx) error {
	// Core process start //
	date := dateStr
	if dateStr == "today" {
		date = time.Now().Format("20060102")
	}
	var index = -1
	if num == "" {
		index = 0
	}

	// Cache
	for i := 0; i < recentSchoolsSize; i++ {
		school := recentSchools[i]
		if strings.Contains(school.SchoolName, decodedSchoolName) && (school.Num == num || num == "") && school.Date == date {
			log.Printf("Cached in %s", school.SchoolName)
			response := school.MenuData
			return c.JSON(response)
		}
	}
	// Cache Done

	// Search school
	schoolDataRes := request.SearchSchool(schoolName)
	if schoolDataRes.Status.Success == false {
		return c.JSON(schoolDataRes)
	}

	searchCnt := len(schoolDataRes.List)
	for i := 0; i < searchCnt && index == -1; i++ {
		if schoolDataRes.List[i].AptCode == num {
			index = i
		}
	}
	if index == -1 {
		schoolDataRes.Status = request.FailResponse("검색된 학교목록에 교육청코드가 '" + num + "'인 학교가 없습니다.")
		return c.JSON(schoolDataRes)
	}
	schoolData := schoolDataRes.List[index]
	// Search school Done

	// Search menu
	menuData := request.SearchMenu(schoolData.AptCode, schoolData.SchoolCode, date)
	menuData.Status.SchoolAptName = schoolData.AptName
	if searchCnt > 1 {
		menuData.Status.Msg += " | 학교가 두 개 이상 검색되었습니다. /[학교명]/" + dateStr + "/[교육청코드] 로 다른학교도 검색해보세요. 교육청코드(apt_code)는 /school/[학교명] 에서 확인할 수 있습니다"
	}
	if menuData.Status.Success == false {
		menuData.Status.SearchedSchool = schoolData.SchoolName
		//return c.JSON(menuData)
	}
	//Search menu Done

	// Cache school data
	recentSchoolsCnt += 1
	if recentSchoolsCnt >= len(recentSchools) {
		recentSchoolsCnt = -1
	}
	var savingSchool RecentSchool
	savingSchool.MenuData = menuData
	savingSchool.SchoolName = schoolData.SchoolName
	savingSchool.AptCode = schoolData.AptCode
	savingSchool.SchoolCode = schoolData.SchoolCode
	savingSchool.Num = num
	savingSchool.Date = date
	recentSchools[recentSchoolsCnt] = savingSchool

	log.Printf("Searched in %s", schoolName)
	return c.JSON(menuData)
}

func main() {
	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:                 true,
		CaseSensitive:           true,
		StrictRouting:           true,
		ServerHeader:            "Fiber",
		AppName:                 "SchoolMenuApi",
		EnableTrustedProxyCheck: true,
	})

	s1 := gocron.NewScheduler(time.Local)
	s1.Every(1).Day().At("23:00").Do(prepareTomorrow)
	s1.Every(1).Day().At("23:59").Do(setTomorrowCache)
	s1.StartAsync()

	// 1분에 100번 요청가능
	app.Use(limiter.New(limiter.Config{
		Max:        30,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			log.Printf(c.IP())
			log.Printf(c.Get("x-forwarded-for"))
			//return c.IP()
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	app.Get("favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	// Cache List
	app.Get("/stat", func(c *fiber.Ctx) error {
		schools := ""
		for i := 0; i < len(recentSchools); i++ {
			schools += recentSchools[i].SchoolName + "\n"
		}
		return c.SendString(schools)
	})

	// School search
	app.Get("/school/:school", func(c *fiber.Ctx) error {
		schoolName := c.Params("school")
		schoolData := request.SearchSchool(schoolName)
		return c.JSON(schoolData)
	})

	app.Get("/:school", func(c *fiber.Ctx) error {
		schoolName := c.Params("school")
		decodedSchoolName, err := url.QueryUnescape(schoolName)
		if err != nil {
			return c.JSON(request.FailResponse("학교이름이 올바른 형식이 아닙니다"))
		}
		date := "today"
		num := ""

		return apiMainProcess(schoolName, decodedSchoolName, date, num, c)
	})

	app.Get("/:school/:date", func(c *fiber.Ctx) error {
		schoolName := c.Params("school")
		decodedSchoolName, err := url.QueryUnescape(schoolName)
		if err != nil {
			return c.JSON(request.FailResponse("학교이름이 올바른 형식이 아닙니다"))
		}
		date := c.Params("date", "today")
		num := ""

		return apiMainProcess(schoolName, decodedSchoolName, date, num, c)
	})

	app.Get("/:school/:date/:num", func(c *fiber.Ctx) error {
		schoolName := c.Params("school")
		decodedSchoolName, err := url.QueryUnescape(schoolName)
		if err != nil {
			return c.JSON(request.FailResponse("학교이름이 올바른 형식이 아닙니다"))
		}
		date := c.Params("date", "today")
		num := c.Params("num")

		return apiMainProcess(schoolName, decodedSchoolName, date, num, c)
	})

	log.Fatal(app.Listen(":3503"))
}
