package adminController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
}

func (Login) Register(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}
