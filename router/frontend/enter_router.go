package frontend

import (
	"eat/utils/response"
	"github.com/gin-gonic/gin"
)

func IndexRouter(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		response.SuccessWithMessage("api", c)
	})
}
