package admin

import "go_work/four_phase/modules"

type User struct {
	modules.BaseModel
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
}

func (User) TableName() string {
	return "user"
}
