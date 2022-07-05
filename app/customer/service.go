package customer

type Service interface {
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindAllCustomer() {

}

func (s *service) FindOneCustomerByID() {

}

func (s *service) CreateCustomer() {

}

func (s *service) UpdateCustomer() {

}

func (s *service) RemoveCustomerByID() {

}

func (s *service) DestoryCustomerByID() {

}
