package routers

import (
	"go_work/four_phase/controllers"
	"go_work/four_phase/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRouters(r *gin.Engine) {
	router := r.Group("/comment")
	{
		router.GET("/list", controllers.CommentController{}.List)
		router.POST("/add", middleware.AuthRequired(), controllers.CommentController{}.Add)

	}
}
