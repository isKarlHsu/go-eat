package service

import (
	"eat/app/model"
)

type userServ struct {
	BaseService // 嵌入基础服务类
}

var UserAuth *model.UserModel

// UserService 创建服务类实例的函数
func UserService() *userServ {
	service := &userServ{}
	service.GetInstance() // 获取基础服务类的实例
	return service
}

