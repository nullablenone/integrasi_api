package user

import "integrasi_api/internal/integration/jsonplaceholder"

type Service interface {
	SyncUsers() error
	GetAllUsers() ([]User, error)
}

type service struct {
	ExternalUserService jsonplaceholder.ExternalUserService
	Repo                Repository
}

func NewUserService(exUserService jsonplaceholder.ExternalUserService, repo Repository) Service {
	return &service{ExternalUserService: exUserService, Repo: repo}
}

func (s *service) SyncUsers() error {
	users, err := s.ExternalUserService.SyncUsersService()
	if err != nil {
		return err
	}

	for _, u := range users {
		mappedUser := User{
			ID:       uint(u.ID),
			Name:     u.Name,
			Username: u.Username,
			Email:    u.Email,
			Phone:    u.Phone,
			Website:  u.Website,
			Address: Address{
				Street:  u.Address.Street,
				Suite:   u.Address.Suite,
				City:    u.Address.City,
				Zipcode: u.Address.Zipcode,
			},
			Company: Company{
				Name:        u.Company.Name,
				CatchPhrase: u.Company.CatchPhrase,
				Bs:          u.Company.Bs,
			},
		}

		// Simpan ke database langsung
		if err := s.Repo.SaveUser(&mappedUser); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil

}
