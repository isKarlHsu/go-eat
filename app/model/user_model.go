package model

type UserModel struct {
	UserId uint `gorm:"primarykey" json:"user_id"`
	Username string `gorm:"size:64" json:"username"`
	Password string `gorm:"size:255" json:"password"`
	UnionId string `gorm:"size:64" json:"union_id"`
	OpenId string `gorm:"size:64" json:"open_id"`
	NickName string `gorm:"size:64" json:"nickname"`
	Avatar string `gorm:"size:1024" json:"avatar"`
	IP string `gorm:"size:20" json:"ip"`
	Timestamp
}

// TableName 自定义表名
func (UserModel) TableName() string {
	return "eat_user"
}
