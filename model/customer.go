package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint64         `json:"id" gorm:"primaryKey;not null"`
	Name      string         `json:"name" gorm:"not null"`
	CitizenID string         `json:"citizen_id" gorm:"not null"`
	PhoneNo   string         `json:"phone_no" gorm:"column:phone_no;not null"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Address   *string        `json:"address" gorm:"column:address"`
	Profile   string         `json:"profile"`
	CreatedAt time.Time      `json:"created_date" gorm:";column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_date" gorm:";column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_date" gorm:"index;column:deleted_at"`
}

func (Customer) TableName() string {
	return "customer"
}
