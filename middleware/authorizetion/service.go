package authorizetion

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	ValidateToken(accessToken string) (jwt.MapClaims, error)
	FindToken(accessToken string) (bool, error)
}

type service struct {
	repository Repository
	secretKey  string
}

func NewService(repository Repository, secretKey string) Service {
	return &service{
		repository: repository,
		secretKey:  secretKey,
	}
}

func (s *service) ValidateToken(accessToken string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(s.secretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New(msgUnauthorize)
	}

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s *service) FindToken(accessToken string) (bool, error) {

	token, err := s.repository.findOneToken(accessToken)

	if err != nil {
		return false, err
	}

	if token != nil {
		return true, nil
	}

	return false, nil
}
