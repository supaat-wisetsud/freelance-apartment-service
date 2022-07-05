package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	ID          uint64         `json:"id" gorm:"primaryKey;not null"`
	User        *Users         `json:"user" gorm:"foreignkey:UserID"`
	UserID      uint64         `json:"user_id" grom:"column:user_id"`
	AccessToken string         `json:"access_token" gorm:"column:access_token;not null"`
	Expired     time.Time      `json:"expired" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_date" gorm:";column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_date" gorm:";column:updated_at;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_date" gorm:"index;column:deleted_at"`
}
