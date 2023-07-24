package frontend

import (
	"eat/app/controller/frontend"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	UserApi := frontend.FrontendApi.UserApi
	router.POST("/mini/userEdit", UserApi.UserEdit)

}
