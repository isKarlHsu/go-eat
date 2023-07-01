package core

import (
	"eat/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitGorm() *gorm.DB {
	mysqlConf := global.Config.Mysql
	if mysqlConf.Host == "" {
		global.Logger.Errorf("未配置mysql host,取消连接")
		return nil
	}
	dsn := mysqlConf.Dsn()
	var mysqlLogger logger.Interface
	// 生产环境只打印错误日志
	if global.Config.System.Env == "release" {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "blog_",
			// SingularTable: false,
			// NameReplacer:  nil,
			// NoLowerCase:   false,
		},
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("[%s] mysql连接失败", dsn))
		panic(any("mysql连接失败" + err.Error()))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns) // 最大空闲连接数
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns) // 最多容纳连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4)       // 连接最大复用数
	return db
}
