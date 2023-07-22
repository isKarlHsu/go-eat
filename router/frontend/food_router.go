package frontend

import (
	"eat/app/controller/frontend"
	"github.com/gin-gonic/gin"
)

func FoodRouter(router *gin.RouterGroup) {
	FoodApi := frontend.FrontendApi.FootApi
	router.GET("/food", FoodApi.Food)
	router.GET("/foods", FoodApi.Foods)
}
