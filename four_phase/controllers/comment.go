package controllers

import (
	"go_work/four_phase/modules/core"
	"go_work/four_phase/modules/posts"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

func (CommentController) Add(c *gin.Context) {
	comment := posts.Comment{}
	c.ShouldBindJSON(&comment)
	err := core.DB.Create(&comment).Error
	code, res := core.Err("comment - Add", err)
	c.JSON(code, res)
}

func (CommentController) List(c *gin.Context) {
	comment := []posts.Comment{}
	err := core.DB.Find(&comment).Error
	code, res := core.DataErr("post - List", err, comment)
	c.JSON(code, res)
}
