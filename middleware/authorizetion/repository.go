package authorizetion

import (
	"apartment/model"
	"errors"

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

	var token model.Token
	if err := r.db.Where("access_token = ?", accessToken).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &token, nil
}
