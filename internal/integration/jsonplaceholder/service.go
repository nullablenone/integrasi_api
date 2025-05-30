package jsonplaceholder

type ExternalUserService interface {
	SyncUsersService() ([]User, error)
}

type externalUserService struct {
	Client JSONPlaceholderClientInterface
}

func NewExternalUserService(client JSONPlaceholderClientInterface) ExternalUserService {
	return &externalUserService{Client: client}
}

func (s *externalUserService) SyncUsersService() ([]User, error) {
	return s.Client.SyncUsers()
}
