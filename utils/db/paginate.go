package db

import (
	"gorm.io/gorm"
)

type Page struct {
	CurrentPage int64 `json:"current_page"`
	PageSize    int64 `json:"page_size"`
	Total       int64 `json:"total"`
	Pages       int64 `json:"pages"`
	List        any `json:"list"`
}

// Paginate 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
