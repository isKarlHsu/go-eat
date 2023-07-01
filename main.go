package main

import (
	"eat/app/command"
	"eat/app/middleware"
	"eat/core"
	"eat/global"
	"eat/router"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化logger配置
	global.Logger = core.InitLogger()
	// 初始化mysql配置
	global.DB = core.InitGorm()
	// 初始化redis
	global.Redis = core.ConnectRedis()
	// 初始化命令行
	command.Execute()

	r := router.InitRouter()
	r.Use(middleware.Cors())
	global.Logger.Infof("服务运行在: %s", global.Config.System.Addr())
	r.Run(global.Config.System.Addr())
}
