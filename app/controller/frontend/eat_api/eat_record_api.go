package eat_api

import (
	"eat/app/model"
	"eat/global"
	"eat/utils/db"
	"eat/utils/response"
	"github.com/gin-gonic/gin"
	"math"
)

type EatRecordParams struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (EatApi) Record (c *gin.Context) {
	var params EatRecordParams
	err := c.ShouldBind(&params)
	if err != nil {
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	global.Logger.Info(params)

	var recordModel []model.RecordModel
	global.DB.Order("record_id Desc").Find(&recordModel)
	count := int64(len(recordModel))
	global.DB.Scopes(db.Paginate(params.Page, params.PageSize)).Order("record_id Desc").Find(&recordModel)
	data := db.Page{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.PageSize),
		Total: 		 count,
		Pages:       int64(math.Ceil(float64(count) / float64(params.PageSize))),
		List:        recordModel,
	}
	response.SuccessWithData(data, c)
}
