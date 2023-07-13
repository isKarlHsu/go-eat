package test

import (
	"eat/core"
	"eat/utils/jwt"
	"testing"
)

func Test_Jwt(t *testing.T) {
	// 读取配置文件
	core.InitConfig()

	user := jwt.JwtPayLoad{
		UserId:   100001,
		Username: "admin",
	}
	token, _ := jwt.GenToken(user)
	t.Log("jwt token生成结果", token)

	user2, _ := jwt.ParseToken(token)
	t.Log("解析结果", user2)
}
