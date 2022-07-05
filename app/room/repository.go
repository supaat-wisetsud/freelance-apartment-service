package room

import (
	"apartment/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindOne(id uint64) (*model.Rooms, error)
	FindAll() ([]model.Rooms, error)
	Create(room *model.Rooms) error
	Update(room *model.Rooms, id uint64) error
	Remove(id uint64) error
	Destroy(id uint64) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindOne(id uint64) (*model.Rooms, error) {
	var room model.Rooms

	if err := r.db.First(&room, id).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *repository) FindAll() ([]model.Rooms, error) {
	var rooms []model.Rooms

	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *repository) Create(room *model.Rooms) error {

	if err := r.db.Create(room).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(room *model.Rooms, id uint64) error {

	if err := r.db.Model(&model.Rooms{}).Where("id = ?", id).Updates(room).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Remove(id uint64) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Rooms{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Destroy(id uint64) error {
	if err := r.db.Unscoped().Where("id = ?", id).Delete(&model.Rooms{}).Error; err != nil {
		return err
	}
	return nil
}
