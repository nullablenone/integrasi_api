package user

import "integrasi_api/internal/integration/jsonplaceholder"

type Service interface {
	GetAllUsersService() ([]jsonplaceholder.User, error)
}

type service struct {
	ExternalUserService jsonplaceholder.ExternalUserService
}

func NewUserService(exUserService jsonplaceholder.ExternalUserService) Service {
	return &service{ExternalUserService: exUserService}
}

func (s *service) GetAllUsersService() ([]jsonplaceholder.User, error) {
	users, err := s.ExternalUserService.FetchUsersService()
	if err != nil {
		return nil, err
	}

	return users, nil
}
