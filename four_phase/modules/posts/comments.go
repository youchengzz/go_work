package posts

import "go_work/four_phase/modules"

type Comment struct {
	modules.BaseModel
	Content string
	UserId  int
	PostId  int
}
