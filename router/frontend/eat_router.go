package frontend

import (
	"eat/app/controller/frontend"
	"github.com/gin-gonic/gin"
)

func EatRouter(router *gin.RouterGroup) {
	EatApi := frontend.FrontendApi.EatApi
	router.POST("/eat/confirm", EatApi.Confirm)
	router.POST("/eat/record", EatApi.Record)
}
