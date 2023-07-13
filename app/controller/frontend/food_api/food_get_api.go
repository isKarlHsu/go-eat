package food_api

import (
	"eat/app/model"
	"eat/global"
	"eat/utils/db"
	"eat/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
)

type FoodsParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Name string `json:"name"`
}

func (FoodApi) Foods(c *gin.Context) {
	var params FoodsParams
	err := c.ShouldBind(&params)
	if err != nil {
		fmt.Println(err)
		response.ErrorWithMessage(err.Error(), c)
		return
	}
	global.Logger.Info(params)
	var foodModel []model.FoodModel
	global.DB.Scopes(model.FoodFilter(params)).Order("food_id Desc").Find(&foodModel)
	count := int64(len(foodModel))
	global.DB.
		Scopes(model.FoodFilter(params)).
		Scopes(db.Paginate(params.Page, params.PageSize)).
		Order("food_id Desc").
		Find(&foodModel)
	data := db.Page{
		CurrentPage: int64(params.Page),
		PageSize:    int64(params.PageSize),
		Total: 		 count,
		Pages:       int64(math.Ceil(float64(count) / float64(params.PageSize))),
		List:        foodModel,
	}
	response.SuccessWithData(data, c)
}