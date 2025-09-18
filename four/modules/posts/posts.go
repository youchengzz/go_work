package posts

import "go_work/four/modules/admin"

type Post struct {
	Id      int
	Title   string
	Content string
	UserId  int
	admin.Base
}
