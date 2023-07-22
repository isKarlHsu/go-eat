package model

type RecordModel struct {
	RecordId uint `gorm:"primarykey" json:"record_id"`
	UserId uint `json:"user_id"`
	Title string `gorm:"size:64" json:"title"`
	Longitude string `gorm:"size:64" json:"longitude"`
	Latitude string `gorm:"size:64" json:"latitude"`
	Timestamp
}

// TableName 自定义表名
func (RecordModel) TableName() string {
	return "eat_record"
}

