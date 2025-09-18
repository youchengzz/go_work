package posts

import "go_work/four/modules/admin"

type Comment struct {
	Id      int
	Content string
	UserId  int
	PostId  int
	admin.Base
}