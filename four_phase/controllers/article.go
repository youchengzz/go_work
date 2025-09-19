package controllers

import (
	"go_work/four_phase/modules/core"
	"go_work/four_phase/modules/posts"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

func (ArticleController) Add(c *gin.Context) {
	post := posts.Post{}
	c.ShouldBindJSON(&post)
	err := core.DB.Create(&post).Error
	code, res := core.Err("post - Add", err)
	c.JSON(code, res)
}

func (ArticleController) Edit(c *gin.Context) {
	post := posts.Post{}
	c.ShouldBindJSON(&post)
	err := core.DB.Create(&post).Error
	code, res := core.Err("post - Edit", err)
	c.JSON(code, res)
}

func (ArticleController) DELETE(c *gin.Context) {
	id := c.Query("id")
	err := core.DB.Where("id = ?", id).Delete(&posts.Post{}).Error
	code, res := core.Err("post - DELETE", err)
	c.JSON(code, res)
}

func (ArticleController) List(c *gin.Context) {
	posts := []posts.Post{}
	err := core.DB.Find(&posts).Error
	code, res := core.DataErr("post - List", err, posts)
	c.JSON(code, res)
}

func (ArticleController) QueryById(c *gin.Context) {
	id := c.Query("id")
	post := posts.Post{}
	err := core.DB.Where("id = ?", id).First(&post).Error
	code, res := core.Err("post - QueryById", err)
	c.JSON(code, res)
}
