package model

type Default struct {
	Success    bool   `json:"success"`
	Msg        string `json:"msg"`
	ServerDate string `json:"server_date"`
	//ResponseTime   string `json:"response_time"`
	SearchedSchool string `json:"searched_school"`
	SchoolAptName  string `json:"school_apt_name"`
}
