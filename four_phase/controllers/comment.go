package controllers

import (
	"go_work/four_phase/modules/admin"
	core "go_work/four_phase/config"
	"go_work/four_phase/modules/posts"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

func (CommentController) Add(c *gin.Context) {
	value, ok := c.Get("user")
	if !ok {
		core.Err500("请先登录", c)
		return
	}
	user := value.(*admin.User)
	comment := posts.Comment{}
	c.ShouldBindJSON(&comment)
	comment.UserId = int(user.BaseModel.Id)
	err := core.DB.Create(&comment).Error
	code, res := core.Err("comment - Add", err)
	c.JSON(code, res)
}

func (CommentController) List(c *gin.Context) {
	postId := c.Query("postId")
	if postId == "" {
		core.Err200("", c)
		return
	}
	comment := []posts.Comment{}
	err := core.DB.Where("post_id = ?", postId).Find(&comment).Error
	code, res := core.DataErr("post - List", err, comment)
	c.JSON(code, res)
}
