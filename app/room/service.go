package room

import (
	"apartment/model"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

type Service interface {
	FindAllRoom() ([]model.Rooms, error)
	FindOneRoomByID(id uint64) (*model.Rooms, error)
	CreateRoom(name string, customerID uint64) error
	UpdateRoom(name string, customerID uint64, active bool, id uint64) error
	RemoveRoomByID(id uint64) error
	DestoryRoomByID(id uint64) error
	UpdatePictureByID(file *multipart.FileHeader, id uint64) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindAllRoom() ([]model.Rooms, error) {

	rooms, err := s.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *service) FindOneRoomByID(id uint64) (*model.Rooms, error) {

	room, err := s.repository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *service) CreateRoom(name string, customerID uint64) error {

	room := model.Rooms{
		Name:       name,
		CustomerID: customerID,
	}

	if err := s.repository.Create(&room); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateRoom(name string, customerID uint64, active bool, id uint64) error {
	room := model.Rooms{
		Name:       name,
		CustomerID: customerID,
		Active:     active,
	}
	if err := s.repository.Update(&room, id); err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveRoomByID(id uint64) error {

	if err := s.repository.Remove(id); err != nil {
		return err
	}

	return nil
}

func (s *service) DestoryRoomByID(id uint64) error {

	if err := s.repository.Destroy(id); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdatePictureByID(file *multipart.FileHeader, id uint64) error {

	path := "/public/images/rooms/" + strconv.Itoa(int(id))
	// Multipart form
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileName := path + "_" + file.Filename
	// Destination
	dst, err := os.Create("." + fileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return err
	}

	room := model.Rooms{
		Picture: fileName,
	}
	if err := s.repository.Update(&room, id); err != nil {
		return err
	}

	return nil
}
