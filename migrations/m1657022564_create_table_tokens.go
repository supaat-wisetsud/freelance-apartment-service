package migrations

import (
	"apartment/model"
	"context"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1657022564CreateTableTokens() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1657022564",
		Migrate: func(db *gorm.DB) error {
			if err := db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
				if err := tx.Migrator().CreateTable(&model.Token{}); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
			return nil
		},
		Rollback: func(db *gorm.DB) error {
			err := db.WithContext(context.Background()).Migrator().DropTable(&model.Token{})
			return err
		},
	}
}
