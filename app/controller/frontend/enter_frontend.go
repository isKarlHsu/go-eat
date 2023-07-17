package frontend

import (
	"eat/app/controller/frontend/food_api"
	"eat/app/controller/frontend/wechat_api"
)

type FrontendApiGroup struct {
	FootApi food_api.FoodApi
	WechatApi wechat_api.WechatApi
}

var FrontendApi = new(FrontendApiGroup)