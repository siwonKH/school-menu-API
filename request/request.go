package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"schoolMenuApi/model"
	"schoolMenuApi/model/apiResponse"
	"schoolMenuApi/model/neisResponse"
	"strconv"
	"strings"
	"time"
)

var KEY = ""

func FailResponse(msg string) model.Default {
	log.Printf("ERROR: %s", msg)
	var response model.Default
	response.Success = false
	response.Msg = "fail | " + msg
	response.ServerDate = time.Now().Format("2006-01-02")
	return response
}

func SuccessResponse(searchedSchool string) model.Default {
	log.Printf("SUCCESS: %s ", searchedSchool)
	var response model.Default
	response.Success = true
	response.Msg = "success"
	response.ServerDate = time.Now().Format("2006-01-02")
	response.SearchedSchool = searchedSchool
	return response
}

func SearchSchool(schoolName string) apiResponse.School {
	var response apiResponse.School
	urlIn := "https://open.neis.go.kr/hub/schoolInfo?Type=json&pSize=10"
	urlIn += "&SCHUL_NM=" + schoolName + "&KEY=" + KEY

	resp, err := http.Get(urlIn)
	defer resp.Body.Close()
	if err != nil {
		response.Status = FailResponse("Neis API에 요청을 실패했습니다")
		return response
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Status = FailResponse("Neis API에서 받은 데이터를 읽지 못했습니다")
		return response
	}

	var schoolData neisResponse.School
	err = json.Unmarshal([]byte(data), &schoolData)
	if err != nil {
		response.Status = FailResponse("Neis API에서 받은 데이터를 읽지 못했습니다")
		return response
	}
	if len(schoolData.SchoolInfo) < 1 {
		response.Status = FailResponse("검색된 학교가 없습니다")
		return response
	}
	info := schoolData.SchoolInfo[1].Row

	response.Status = SuccessResponse(info[0].SchulNm)
	for i := 0; i < len(info) && i < 10; i++ {
		var list apiResponse.SchoolList
		list.AptCode = info[i].AtptOfcdcScCode
		list.AptName = info[i].AtptOfcdcScNm
		list.SchoolCode = info[i].SdSchulCode
		list.SchoolName = info[i].SchulNm
		response.List = append(response.List, list)
	}
	return response
}

func stripNums(menuList []string) []string {
	for i := 0; i < len(menuList); i++ {
		if strings.Contains(menuList[i], ".") {
			menuList[i] = menuList[i][:strings.Index(menuList[i], ".")-1]
			menuList[i] = strings.Replace(menuList[i], "1", "", -1)
		}
	}
	return menuList
}

func SearchMenu(aptCode string, schoolCode string, date string) apiResponse.Menu {
	var response apiResponse.Menu
	if date == "today" {
		date = time.Now().Format("20060102")
	}
	response.MenuDate = date
	urlIn := "https://open.neis.go.kr/hub/mealServiceDietInfo?Type=json" + "&KEY=" + KEY
	urlIn += "&ATPT_OFCDC_SC_CODE=" + aptCode + "&SD_SCHUL_CODE=" + schoolCode + "&MLSV_YMD=" + date

	resp, err := http.Get(urlIn)
	defer resp.Body.Close()
	if err != nil {
		response.Status = FailResponse("Neis API에 요청을 실패했습니다")
		return response
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Status = FailResponse("Neis API에서 받은 데이터를 읽지 못했습니다")
		return response
	}

	var neisMenuData neisResponse.Menu
	err = json.Unmarshal([]byte(data), &neisMenuData)
	if err != nil {
		response.Status = FailResponse("Neis API에서 받은 데이터를 읽지 못했습니다")
		return response
	}
	if neisMenuData.MealServiceDietInfo == nil {
		response.Status = FailResponse("오늘의 학교 급식이 존재하지 않습니다")
		return response
	}

	menuData := neisMenuData.MealServiceDietInfo[1].Row
	var menuList [4][]string
	for i := 0; i < len(menuData); i++ {
		menuType, err := strconv.Atoi(menuData[i].MmealScCode)
		if err != nil {
			response.Status = FailResponse("Neis API 구조 변경됨. 관리자에게 문의 바람")
			return response
		}
		menuList[menuType] = strings.Split(menuData[i].DdishNm, "<br/>")
	}

	response.Status = SuccessResponse(menuData[0].SchulNm)
	response.Breakfast = stripNums(menuList[1])
	response.Lunch = stripNums(menuList[2])
	response.Dinner = stripNums(menuList[3])
	return response
}
