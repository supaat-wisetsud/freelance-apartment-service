package model

import (
	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	ID         uint64         `json:"id" gorm:"primaryKey;not null"`
	Username   string         `json:"username" gorm:"not null;unique"`
	Password   string         `json:"password" gorm:"not null"`
	Name       string         `json:"name" gorm:"not null;unique"`
	PhoneNo    *string        `json:"phone_no" gorm:"column:phone_no;unique"`
	Email      *string        `json:"email" gorm:"unique"`
	MaxSession int            `json:"max_session" gorm:"coulmn:max_session;default:10"`
	Enable     bool           `json:"enable" gorm:"default:true"`
	CreatedAt  time.Time      `json:"created_date" gorm:";column:created_at;autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_date" gorm:";column:updated_at;autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_date" gorm:"index;column:deleted_at"`
}

func (Users) TableName() string {
	return "users"
}

func (u *Users) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	if err != nil {
		return err
	}
	u.Password = base64.StdEncoding.EncodeToString(hash)
	return nil
}

func (u *Users) ComparePassword(password string) error {
	hashedPassword, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}
