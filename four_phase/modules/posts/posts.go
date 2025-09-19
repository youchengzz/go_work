package posts

import "go_work/four_phase/modules"

type Post struct {
	modules.BaseModel
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int
}
