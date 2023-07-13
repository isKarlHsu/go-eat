package model

import (
	"gorm.io/gorm"
	"reflect"
)

type FoodModel struct {
	FoodId uint `gorm:"primarykey" json:"food_id"`
	Name string `gorm:"size:64" json:"name"`
	Timestamp
}

// TableName 自定义表名
func (FoodModel) TableName() string {
	return "eat_food"
}

func FoodFilter(params any) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		immutableV := reflect.ValueOf(params)
		name := immutableV.FieldByName("Name").String()
		if name != "" {
			db.Where("name like ?", "%"+name+"%")
		}
		return db
	}
}