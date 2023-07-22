package eat_api

import (
	"eat/app/model"
	"eat/app/service"
	"eat/global"
	"eat/utils/response"
	"github.com/gin-gonic/gin"
)

type EatConfirmParams struct {
	Title string `json:"title"`
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
}

func (EatApi) Confirm (c *gin.Context) {
	var params EatConfirmParams
	err := c.ShouldBind(&params)
	if err != nil {
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	global.Logger.Info(params)

	var recordModel model.RecordModel
	recordModel.UserId = service.UserAuth.UserId
	recordModel.Title = params.Title
	recordModel.Longitude = params.Longitude
	recordModel.Latitude = params.Latitude
	global.DB.Save(&recordModel)

	response.SuccessWithMessage("保存成功", c)
}
