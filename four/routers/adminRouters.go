package routers

import (
	"go_work/four/controllers/adminController"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	router := r.Group("/admin")
	{
		router.POST("/register", adminController.Login{}.Register)
	}
}
