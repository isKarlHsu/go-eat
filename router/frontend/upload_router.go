package frontend

import (
	"eat/app/controller/frontend"
	"github.com/gin-gonic/gin"
)

func UploadRouter(router *gin.RouterGroup) {
	UploadApi := frontend.FrontendApi.UploadApi
	router.POST("/upload/image", UploadApi.Image)
}
