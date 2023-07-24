package user_api

import (
	"eat/app/model"
	"eat/app/service"
	"eat/global"
	"eat/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserEditParams struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (UserApi) UserEdit(c *gin.Context) {
	var params UserEditParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	global.Logger.Info(params)
	var userModel model.UserModel
	global.DB.Where("user_id = ?", service.UserAuth.UserId).First(&userModel)
	userModel.Nickname = params.Nickname
	userModel.Avatar = params.Avatar
	global.DB.Save(&userModel)

	response.SuccessWithMessage("保存成功", c)
}
