package neisResponse

type School struct {
	SchoolInfo []struct {
		Head []struct {
			ListTotalCount int `json:"list_total_count,omitempty"`
			Result         struct {
				Code    string `json:"CODE"`
				Message string `json:"MESSAGE"`
			} `json:"RESULT,omitempty"`
		} `json:"head,omitempty"`
		Row []struct {
			AtptOfcdcScCode       string      `json:"ATPT_OFCDC_SC_CODE"`
			AtptOfcdcScNm         string      `json:"ATPT_OFCDC_SC_NM"`
			SdSchulCode           string      `json:"SD_SCHUL_CODE"`
			SchulNm               string      `json:"SCHUL_NM"`
			EngSchulNm            string      `json:"ENG_SCHUL_NM"`
			SchulKndScNm          string      `json:"SCHUL_KND_SC_NM"`
			LctnScNm              string      `json:"LCTN_SC_NM"`
			JuOrgNm               string      `json:"JU_ORG_NM"`
			FondScNm              string      `json:"FOND_SC_NM"`
			OrgRdnzc              string      `json:"ORG_RDNZC"`
			OrgRdnma              string      `json:"ORG_RDNMA"`
			OrgRdnda              string      `json:"ORG_RDNDA"`
			OrgTelno              string      `json:"ORG_TELNO"`
			HmpgAdres             string      `json:"HMPG_ADRES"`
			CoeduScNm             string      `json:"COEDU_SC_NM"`
			OrgFaxno              string      `json:"ORG_FAXNO"`
			HsScNm                string      `json:"HS_SC_NM"`
			IndstSpeclCccclExstYn string      `json:"INDST_SPECL_CCCCL_EXST_YN"`
			HsGnrlBusnsScNm       string      `json:"HS_GNRL_BUSNS_SC_NM"`
			SpclyPurpsHsOrdNm     interface{} `json:"SPCLY_PURPS_HS_ORD_NM"`
			EneBfeSehfScNm        string      `json:"ENE_BFE_SEHF_SC_NM"`
			DghtScNm              string      `json:"DGHT_SC_NM"`
			FondYmd               string      `json:"FOND_YMD"`
			FoasMemrd             string      `json:"FOAS_MEMRD"`
			LoadDtm               string      `json:"LOAD_DTM"`
		} `json:"row,omitempty"`
	} `json:"schoolInfo"`
}
