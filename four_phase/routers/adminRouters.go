package routers

import (
	"go_work/four_phase/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	router := r.Group("/belogs")
	{
		router.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		router.POST("/register", controllers.LoginController{}.Register)
		router.POST("/login", controllers.LoginController{}.Login)

		// router.GET("/user", adminController.Login{}.UserInfo)
	}
}
