package food_api

import (
	"eat/app/model"
	"eat/app/service"
	"eat/global"
	"eat/utils/db"
	"eat/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
)

func (FoodApi) Food(c *gin.Context) {
	foodService := service.FoodService()
	var result model.FoodModel
	for {
		food, count := foodService.GetFoodForRand()
		if count > 0 {
			result = food
			break
		}
	}
	// foodService2 := service.FoodService()
	// foodService2.GetFoodForRand()

	response.SuccessWithData(result, c)
}

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