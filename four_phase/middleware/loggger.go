package middleware

import (
	"encoding/json"
	core "go_work/four_phase/config"
	"go_work/four_phase/modules/admin"
	"go_work/four_phase/modules/logs"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logs.Log{
			Path: c.Request.URL.Path,
			Ip:   c.ClientIP(),
		}
		token := c.GetHeader("AccessToken")
		var claims *admin.CustomClaims
		if token != "" {
			claims, _ = core.ParseToken(token)
			log.UserId = int(claims.Id)
		}
		c.Next()
		log.Status = c.Writer.Status()
		log.ErrorMessage = c.Errors.String()
		if c.Request.Method == "POST" {
			bytes, _ := c.GetRawData()
			log.RequartParam = string(bytes)
		} else {
			queryParams := c.Request.URL.Query()
			if len(queryParams) != 0 {
				params := make(map[string]string, 0)
				for k, v := range queryParams {
					params[k] = v[0]
				}
				jsonByte, err := json.Marshal(params)
				if err == nil {
					log.RequartParam = string(jsonByte)
				}
			}
		}
		core.DB.Table("logs").Create(&log)
	}
}
