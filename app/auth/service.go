package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	generateAccessToken(username string, password string) (string, error)
	registerUser(username string, password string, name string, phoneNo string, email string) error
	revoke(accessToken string) error
}

type service struct {
	secretKey  string
	repository Repository
}

func NewService(repository Repository, secretKey string) Service {
	return &service{
		repository: repository,
		secretKey:  secretKey,
	}
}

func (s *service) registerUser(username string, password string, name string, phoneNo string, email string) error {
	if err := s.repository.createUser(username, password, name, email, phoneNo); err != nil {
		return err
	}
	return nil
}

func (s *service) generateAccessToken(username string, password string) (string, error) {

	user, err := s.repository.findUserByUsername(username)

	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New(mgsUserNotFound)
	}

	if err := user.ComparePassword(password); err != nil {
		return "", err
	}

	countToken, err := s.repository.countAccessToken(user.ID)
	if err != nil {
		return "", err
	}

	if countToken >= int64(user.MaxSession) {
		return "", errors.New(msgSessionLimit)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
	})

	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}

	if err := s.repository.addAccessToken(t, user.ID); err != nil {
		return "", err
	}

	return t, nil
}

func (s *service) revoke(accessToken string) error {

	if err := s.repository.removeAccessToken(accessToken); err != nil {
		return err
	}

	return nil
}
