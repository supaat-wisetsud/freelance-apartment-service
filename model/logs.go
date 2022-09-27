package model

import (
	"time"

	"gorm.io/gorm"
)

type Logs struct {
	ID         uint64         `json:"id" gorm:"primaryKey;not null"`
	RoomsID    uint64         `json:"room_id" gorm:"column:room_id;not null"`
	Rooms      Rooms          `json:"room"`
	CustomerID uint64         `json:"customer_id" gorm:"column:customer_id;not null"`
	Customer   Customer       `json:"customer"`
	Note       string         `json:"note" gorm:"column:note"`
	CreatedAt  time.Time      `json:"created_date" gorm:";column:created_at;autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_date" gorm:";column:updated_at;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_date" gorm:"index;column:deleted_at"`
}

func (Logs) TableName() string {
	return "logs"
}
