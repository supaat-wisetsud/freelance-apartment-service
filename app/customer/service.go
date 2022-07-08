package customer

import (
	"apartment/model"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

type Service interface {
	FindAllCustomer() ([]model.Customer, error)
	FindOneCustomerByID(id uint64) (*model.Customer, error)
	CreateCustomer(name string, citizenID string, phoneNo string, email string, address string) error
	UpdateCustomer(name string, citizenID string, phoneNo string, email string, address string, id uint64) error
	RemoveCustomerByID(id uint64) error
	DestoryCustomerByID(id uint64) error
	UpdateProfileByID(file *multipart.FileHeader, id uint64) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindAllCustomer() ([]model.Customer, error) {

	customers, err := s.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *service) FindOneCustomerByID(id uint64) (*model.Customer, error) {

	customer, err := s.repository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *service) CreateCustomer(name string, citizenID string, phoneNo string, email string, address string) error {

	customer := model.Customer{
		Name:      name,
		CitizenID: citizenID,
		PhoneNo:   phoneNo,
		Email:     email,
		Address:   &address,
	}

	if err := s.repository.Create(&customer); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCustomer(name string, citizenID string, phoneNo string, email string, address string, id uint64) error {
	customer := model.Customer{
		Name:      name,
		CitizenID: citizenID,
		PhoneNo:   phoneNo,
		Email:     email,
		Address:   &address,
	}
	if err := s.repository.Update(&customer, id); err != nil {
		return err
	}

	return nil
}

func (s *service) RemoveCustomerByID(id uint64) error {

	if err := s.repository.Remove(id); err != nil {
		return err
	}

	return nil
}

func (s *service) DestoryCustomerByID(id uint64) error {

	if err := s.repository.Destroy(id); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateProfileByID(file *multipart.FileHeader, id uint64) error {

	dir := "/public/images/customers"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	path := dir + "/" + strconv.Itoa(int(id))
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

	customer := model.Customer{
		Profile: fileName,
	}
	if err := s.repository.Update(&customer, id); err != nil {
		return err
	}

	return nil
}
