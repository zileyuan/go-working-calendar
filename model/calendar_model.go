package model

type CalendarModel struct {
	Maintain []string `json:"maintain"`
	Working  []string `json:"working"`
	Holiday  []string `json:"holiday"`
}
