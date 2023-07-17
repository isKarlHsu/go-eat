package frontend

import (
	"eat/app/controller/frontend"
	"github.com/gin-gonic/gin"
)

func WechatRouter(router *gin.RouterGroup) {
	WechatApi := frontend.FrontendApi.WechatApi
	router.POST("/mini/user", WechatApi.User)
}
