package customer

import (
	"apartment/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindOne(id uint64) (*model.Customer, error)
	FindAll() ([]model.Customer, error)
	Create(customer *model.Customer) error
	Update(customer *model.Customer, id uint64) error
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

func (r *repository) FindOne(id uint64) (*model.Customer, error) {
	var customer model.Customer

	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *repository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer

	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *repository) Create(customer *model.Customer) error {

	if err := r.db.Create(customer).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(customer *model.Customer, id uint64) error {

	if err := r.db.Model(customer).Update("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Remove(id uint64) error {
	if err := r.db.Where("id = ?").Delete(&model.Customer{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Destroy(id uint64) error {
	if err := r.db.Unscoped().Where("id = ?").Delete(&model.Customer{}).Error; err != nil {
		return err
	}
	return nil
}
