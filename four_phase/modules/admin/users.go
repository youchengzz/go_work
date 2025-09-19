package admin

import "go_work/four_phase/modules"

type User struct {
	modules.BaseModel
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

func (User) TableName() string {
	return "user"
}
