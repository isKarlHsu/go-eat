package frontend

import "eat/app/controller/frontend/food_api"

type FrontendApiGroup struct {
	FootApi food_api.FoodApi
}

var FrontendApi = new(FrontendApiGroup)