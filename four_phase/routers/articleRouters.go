package routers

import (
	"go_work/four_phase/controllers"
	"go_work/four_phase/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRouters(r *gin.Engine) {
	router := r.Group("/article")
	{
		router.GET("/list", controllers.ArticleController{}.List)
		router.GET("/queryById", controllers.ArticleController{}.QueryById)
		router.POST("/add", middleware.AuthRequired(), controllers.ArticleController{}.Add)
		router.POST("/edit", middleware.AuthRequired(), controllers.ArticleController{}.Edit)
		router.DELETE("/delete", middleware.AuthRequired(), controllers.ArticleController{}.DELETE)

	}
}
