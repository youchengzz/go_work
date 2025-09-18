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

func ERR(step string, err error) Result {
	result := Result{}
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Success = false
		result.Message = fmt.Sprint("处理异常", step, err)
		return result
	}
	result.Code = http.StatusOK
	result.Success = true
	return result
}
