package logs

import "go_work/four_phase/modules"

type Log struct {
	modules.BaseModel
	Path         string
	RequartParam string
	Status       int
	ErrorMessage string
	UserId       int
	Ip           string
}
