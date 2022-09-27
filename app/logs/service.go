package logs

import (
	"apartment/model"
	"time"
)

type Service interface {
	FindAllLogs() ([]model.Logs, error)
	CreateLogs(customerID uint64, roomID uint64) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindAllLogs() ([]model.Logs, error) {
	return s.repository.FindAll()
}

func (s *service) CreateLogs(customerID uint64, roomID uint64) error {
	log := &model.Logs{
		CustomerID: customerID,
		RoomsID:    roomID,
		Note:       "ย้ายออก",
		CreatedAt:  time.Now(),
	}
	return s.repository.Create(log)
}
