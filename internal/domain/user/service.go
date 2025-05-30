package user

import "integrasi_api/internal/integration/jsonplaceholder"

type Service interface {
	SyncUsers() ([]jsonplaceholder.User, error)
}

type service struct {
	ExternalUserService jsonplaceholder.ExternalUserService
}

func NewUserService(exUserService jsonplaceholder.ExternalUserService) Service {
	return &service{ExternalUserService: exUserService}
}

func (s *service) SyncUsers() ([]jsonplaceholder.User, error) {
	users, err := s.ExternalUserService.SyncUsersService()
	if err != nil {
		return nil, err
	}

	return users, nil
}
