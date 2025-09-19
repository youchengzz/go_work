package config

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int
	Success bool
	Message string
	Data    interface{}
}

func DataErr(step string, err error, data interface{}) (int, Result) {
	result := Result{}
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Success = false
		result.Message = fmt.Sprint("请求处理异常 ", step, err)
		return http.StatusInternalServerError, result
	}
	result.Code = http.StatusOK
	result.Success = true
	result.Data = data
	return http.StatusOK, result
}

func Err(step string, err error) (int, Result) {
	result := Result{}
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Success = false
		result.Message = fmt.Sprint("请求处理异常 ", step, err)
		return http.StatusInternalServerError, result
	}
	result.Code = http.StatusOK
	result.Success = true
	result.Message = "操作成功"
	return http.StatusOK, result
}

func ReturnErr(step string, err error, c *gin.Context) {
	result := Result{}
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Success = false
		result.Message = fmt.Sprint("请求处理异常 ", step, err)
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result.Code = http.StatusOK
	result.Success = true
	c.JSON(http.StatusOK, result)
}

func Err500(msg string, c *gin.Context) {
	result := Result{
		Code:    http.StatusInternalServerError,
		Success: false,
		Message: msg,
	}
	c.JSON(http.StatusInternalServerError, result)
}

func Err403(msg string, c *gin.Context) {
	result := Result{
		Code:    http.StatusForbidden,
		Success: false,
		Message: msg,
	}
	c.JSON(http.StatusForbidden, result)
}

func Err200(msg string, c *gin.Context) {
	result := Result{
		Code:    http.StatusOK,
		Success: false,
		Message: msg,
	}
	c.JSON(http.StatusOK, result)
}
