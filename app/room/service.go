package room

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

func (s *service) FindAllRoom() {

}

func (s *service) FindOneRoomByID() {

}

func (s *service) CreateRoom() {

}

func (s *service) UpdateRoom() {

}

func (s *service) RemoveRoomByID() {

}

func (s *service) DestoryRoomByID() {

}
