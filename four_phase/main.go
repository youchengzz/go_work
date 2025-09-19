package main

import (
	"go_work/four_phase/middleware"
	"go_work/four_phase/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(middleware.Logger())

	routers.AdminRouters(r)
	routers.ArticleRouters(r)
	routers.CommentRouters(r)

	r.Run()
}
