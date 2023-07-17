package router

import (
	"eat/global"
	"eat/router/backend"
	"eat/router/frontend"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 访问静态资源
	router.Static("/resource", "resource")

	// 后台接口组
	backendRouterGroup := router.Group("admin")
	backend.IndexRouter(backendRouterGroup)

	// 前台接口组
	frontendRouterGroup := router.Group("api")
	frontend.IndexRouter(frontendRouterGroup)
	frontend.FoodRouter(frontendRouterGroup)
	frontend.WechatRouter(frontendRouterGroup)
	return router
}
