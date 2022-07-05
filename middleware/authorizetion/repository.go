package authorizetion

import (
	"apartment/model"

	"gorm.io/gorm"
)

type Repository interface {
	findOneToken(accessToken string) (*model.Token, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) findOneToken(accessToken string) (*model.Token, error) {

	r.db.Where("access_token = ?", accessToken).First(&model.Token{})

	return nil, nil
}
