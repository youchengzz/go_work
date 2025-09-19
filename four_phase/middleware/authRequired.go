package middleware

import (
	"go_work/four_phase/modules"
	"go_work/four_phase/modules/admin"
	core "go_work/four_phase/config"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("AccessToken")
		claims, err := core.ParseToken(token)
		if err != nil {
			return
		}
		user := admin.User{
			Username: claims.Username,
			BaseModel: modules.BaseModel{
				Id: claims.Id,
			},
		}
		c.Set("user", &user)
	}
}
