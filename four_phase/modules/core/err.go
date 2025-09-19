package core

import (
	"fmt"
	"net/http"
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
	return http.StatusOK, result
}
