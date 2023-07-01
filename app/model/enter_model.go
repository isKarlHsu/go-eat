package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Timestamp 自定义时间格式
type Timestamp struct {
	CreatedAt *LocalTime `json:"created_at" gorm:"type:datetime"`
	UpdatedAt *LocalTime `json:"updated_at" gorm:"type:datetime"`
}

type LocalTime struct {
	time.Time
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	// 判断给定时间是否和默认零时间的时间戳相同
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Format("2006-01-02 15:04:05"), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *LocalTime) String() string {
	return fmt.Sprintf("hhh:%s", t.String())
}