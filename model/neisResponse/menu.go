package neisResponse

type Menu struct {
	MealServiceDietInfo []struct {
		Head []struct {
			ListTotalCount int `json:"list_total_count,omitempty"`
			Result         struct {
				Code    string `json:"CODE"`
				Message string `json:"MESSAGE"`
			} `json:"RESULT,omitempty"`
		} `json:"head,omitempty"`
		Row []struct {
			AtptOfcdcScCode string `json:"ATPT_OFCDC_SC_CODE"`
			AtptOfcdcScNm   string `json:"ATPT_OFCDC_SC_NM"`
			SdSchulCode     string `json:"SD_SCHUL_CODE"`
			SchulNm         string `json:"SCHUL_NM"`
			MmealScCode     string `json:"MMEAL_SC_CODE"`
			MmealScNm       string `json:"MMEAL_SC_NM"`
			MlsvYmd         string `json:"MLSV_YMD"`
			MlsvFgr         string `json:"MLSV_FGR"`
			DdishNm         string `json:"DDISH_NM"`
			OrplcInfo       string `json:"ORPLC_INFO"`
			CalInfo         string `json:"CAL_INFO"`
			NtrInfo         string `json:"NTR_INFO"`
			MlsvFromYmd     string `json:"MLSV_FROM_YMD"`
			MlsvToYmd       string `json:"MLSV_TO_YMD"`
		} `json:"row,omitempty"`
	} `json:"mealServiceDietInfo"`
}
