package apiResponse

import (
	"schoolMenuApi/model"
)

type Menu struct {
	Status    model.Default `json:"status"`
	MenuDate  string        `json:"menu_date"`
	Breakfast []string      `json:"breakfast"`
	Lunch     []string      `json:"lunch"`
	Dinner    []string      `json:"dinner"`
}
