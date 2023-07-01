package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

const (
	ErrorCode   = 100
	SuccessCode = 101
)

func Result(code int, data any, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}
func Success(data any, message string, c *gin.Context) {
	Result(SuccessCode, data, message, c)
}

func SuccessWithData(data any, c *gin.Context) {
	Result(SuccessCode, data, "success", c)
}

func SuccessWithMessage(message string, c *gin.Context) {
	Result(SuccessCode, map[string]any{}, message, c)
}

func Error(data any, message string, c *gin.Context) {
	Result(ErrorCode, data, message, c)
}

func ErrorWithData(data any, c *gin.Context) {
	Result(ErrorCode, data, "success", c)
}

func ErrorWithMessage(message string, c *gin.Context) {
	Result(ErrorCode, map[string]any{}, message, c)
}
