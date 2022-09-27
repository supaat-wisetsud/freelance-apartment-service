package migrations

import (
	"apartment/model"
	"context"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1664297317CreateTableLogs() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1664297317",
		Migrate: func(db *gorm.DB) error {
			if err := db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
				if err := tx.Migrator().CreateTable(&model.Logs{}); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(db *gorm.DB) error {
			err := db.WithContext(context.Background()).Migrator().DropTable(&model.Logs{})
			return err
		},
	}
}
