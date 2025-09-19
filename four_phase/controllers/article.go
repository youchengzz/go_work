package controllers

import (
	"errors"
	core "go_work/four_phase/config"
	"go_work/four_phase/modules/admin"
	"go_work/four_phase/modules/posts"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

func (ArticleController) Add(c *gin.Context) {
	value, ok := c.Get("user")
	if !ok {
		core.Err500("请先登录", c)
		return
	}
	user := value.(*admin.User)
	post := posts.Post{}
	c.ShouldBindJSON(&post)
	if post.Title == "" || post.Content == "" {
		core.ReturnErr("add post", errors.New("标题或内容不能为空"), c)
		return
	}
	post.UserId = int(user.Id)
	err := core.DB.Create(&post).Error
	code, res := core.Err("post - Add", err)
	c.JSON(code, res)
}

func (ArticleController) Edit(c *gin.Context) {
	value, ok := c.Get("user")
	if !ok {
		core.Err500("请先登录", c)
		return
	}
	user := value.(*admin.User)
	post := posts.Post{}
	c.ShouldBindJSON(&post)
	if post.UserId != int(user.BaseModel.Id) {
		core.Err403("无操作权限", c)
		return
	}
	err := core.DB.Create(&post).Error
	code, res := core.Err("post - Edit", err)
	c.JSON(code, res)
}

func (ArticleController) DELETE(c *gin.Context) {
	value, ok := c.Get("user")
	if !ok {
		core.Err500("请先登录", c)
		return
	}
	user := value.(*admin.User)
	id := c.Query("id")
	post := posts.Post{}
	if err := core.DB.Where("id = ?", id).First(&post).Error; err != nil {
		core.Err500("删除文章失败", c)
		return
	}
	if post.UserId != int(user.BaseModel.Id) {
		core.Err403("无操作权限", c)
		return
	}
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
