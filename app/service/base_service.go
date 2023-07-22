package service

import (
	"eat/global"
	"reflect"
)

var containers = make(map[string]interface{})

// BaseService 基础服务层
type BaseService struct{}

// GetInstance 获取服务实例
func (s *BaseService) GetInstance() interface{} {
	global.Logger.Info("in")
	className := getTypeName(s)
	// 判断是否存在
	if instance, exists := containers[className]; exists {
		return instance
	}
	global.Logger.Info("new")
	// 不存在则实例化
	instance := new(BaseService)
	containers[className] = instance
	return instance
}

// getTypeName 获取类型名称
func getTypeName(i interface{}) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.PkgPath() + "." + t.Name()
}
