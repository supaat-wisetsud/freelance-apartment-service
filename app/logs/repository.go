package logs

import (
	"apartment/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Logs, error)
	Create(logs *model.Logs) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll() ([]model.Logs, error) {
	result := []model.Logs{}
	if err := r.db.Preload("Customer").Preload("Rooms").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Create(logs *model.Logs) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		roomID := (*logs).RoomsID
		if err := r.db.Model(&model.Rooms{}).Where("id = ?", roomID).Update("customer_id", nil).Error; err != nil {
			return err
		}
		if err := r.db.Create(logs).Error; err != nil {
			return err
		}
		return nil
	})
}
