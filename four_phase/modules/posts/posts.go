package posts

import "go_work/four_phase/modules"

type Post struct {
	modules.BaseModel
	Title   string
	Content string
	UserId  int
}
