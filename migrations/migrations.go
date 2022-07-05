package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		m1657022541CreateTableUsers(),
		m1657022564CreateTableTokens(),
		m1657022597CreateTableCustomers(),
		m1657022580CreateTableRooms(),
	})

	return m
}
