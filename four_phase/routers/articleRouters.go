package routers

import (
	"go_work/four_phase/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleRouters(r *gin.Engine) {
	router := r.Group("/article")
	{
		router.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "article.html", gin.H{})
		})

		router.GET("/list", controllers.ArticleController{}.List)
		router.GET("/queryById", controllers.ArticleController{}.QueryById)
		router.POST("/add", controllers.ArticleController{}.Add)
		router.POST("/edit", controllers.ArticleController{}.Edit)
		router.DELETE("/delete", controllers.ArticleController{}.DELETE)
		
	}
}