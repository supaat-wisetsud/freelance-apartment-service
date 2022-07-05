package model

import (
	"time"

	"gorm.io/gorm"
)

type Rooms struct {
	ID         uint64         `json:"id" gorm:"primaryKey;not null"`
	Name       string         `json:"name" gorm:"not null;unique"`
	Active     bool           `json:"active" gorm:"not null;default:false"`
	Customer   *Customer      `json:"customer" gorm:"foreignkey:CustomerID"`
	CustomerID uint64         `json:"customer_id" gorm:"column:customer_id"`
	Picture    string         `json:"picture" sql:"type:json" gorm:"type:json"`
	CreatedAt  time.Time      `json:"created_date" gorm:";column:created_at;autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_date" gorm:";column:updated_at;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_date" gorm:"index;column:deleted_at"`
}

func (Rooms) TableName() string {
	return "rooms"
}
