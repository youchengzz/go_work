package modules

import "time"

type BaseModel struct {
	Id       uint      `gorm:"primarykey"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
}
