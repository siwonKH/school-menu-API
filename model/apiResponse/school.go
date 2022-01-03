package apiResponse

import (
	"schoolMenuApi/model"
)

type School struct {
	Status model.Default `json:"status"`
	List   []SchoolList  `json:"list"`
}
