package middleware

import (
	"eat/app/model"
	"eat/app/service"
	"eat/global"
	"eat/utils/jwt"
	"eat/utils/response"
	"github.com/gin-gonic/gin"
)

func UserJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		tokenStr := c.Request.Header.Get("token")
		// tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuaWNrbmFtZSI6ImtrIiwiYXZhdGFyIjoiaHR0cHM6Ly93d3cua2FybHh1LmNvbS9hZG1pbi9pbWcvcHJvZmlsZV9zbWFsbC5qcGciLCJleHAiOjE2ODk3NjY4MzcuNzEzMzkzLCJpc3MiOiJlYXQifQ.BemTj-rWdHrRVrsswqeKpu8MNWstqM_x1S34de0u5PA"
		global.Logger.Info("token", tokenStr)
		user, err := jwt.ParseToken(tokenStr)
		if err != nil {
			response.Fail(c)
		}
		var userModel model.UserModel
		result := global.DB.Where("user_id = ?", user.UserId).First(&userModel)
		if result.RowsAffected == 0 {
			response.Fail(c)
		}
		// 解析好存储
		service.UserAuth = &userModel
		global.Logger.Info("到这里了")
		// 处理请求
		c.Next()
	}
}