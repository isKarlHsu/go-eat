package service

import (
	"eat/app/model"
	"eat/global"
	"math/rand"
)

type foodServ struct {
	BaseService // 嵌入基础服务类
}

// FoodService 创建服务类实例的函数
func FoodService() *foodServ {
	service := &foodServ{}
	service.GetInstance() // 获取基础服务类的实例
	return service
}

func (*foodServ) GetFoodForRand() (model.FoodModel, int64) {
	var foodModel model.FoodModel
	count := new(int64)
	global.DB.Model(&foodModel).Count(count)
	r := rand.Intn(int(*count))
	result := global.DB.Where("food_id = ?", r).First(&foodModel)
	return foodModel, result.RowsAffected
}
