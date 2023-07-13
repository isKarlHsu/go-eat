package main

import (
	"eat/app/command"
	"eat/app/middleware"
	"eat/core"
	"eat/global"
	"eat/router"
	"flag"
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

	// 判断如果输入了命令则执行命令行工具
	flag.Parse()
	optionNum := flag.NArg()
	if optionNum > 0 {
		// 初始化命令行
		command.Execute()
		return
	}

	// 初始化路由
	r := router.InitRouter()
	// 增加跨域中间件
	r.Use(middleware.Cors())
	global.Logger.Infof("服务运行在: %s", global.Config.System.Addr())
	// 启动http服务
	r.Run(global.Config.System.Addr())
}
