package auth

import (
	"apartment/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	findUserByUsername(username string) (*model.Users, error)
	removeAccessToken(accessToken string) error
	createUser(username string, password string, name string, email string, phoneNo string) error
	addAccessToken(accessToken string, userID uint64) error
	countAccessToken(userID uint64) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) createUser(username string, password string, name string, email string, phoneNo string) error {

	user := model.Users{
		Username: username,
		Name:     name,
		Email:    &email,
		PhoneNo:  &phoneNo,
	}

	user.SetPassword(password)

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) findUserByUsername(username string) (*model.Users, error) {

	var user model.Users
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *repository) addAccessToken(accessToken string, userID uint64) error {

	token := model.Token{
		UserID:      userID,
		AccessToken: accessToken,
		Expired:     time.Now().AddDate(0, 0, 14),
	}

	if err := r.db.Create(&token).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) removeAccessToken(accessToken string) error {

	if err := r.db.Where("access_token = ?", accessToken).Delete(&model.Token{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) countAccessToken(userID uint64) (int64, error) {

	var count int64
	if err := r.db.Table("tokens").Where("expired >= ? AND user_id = ? AND deleted_at IS NULL", time.Now(), userID).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
