package main

import (
	"go_work/four_phase/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	routers.AdminRouters(r)
	routers.ArticleRouters(r)

	r.Run()
}
