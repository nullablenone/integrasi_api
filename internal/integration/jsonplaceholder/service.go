package jsonplaceholder

type ExternalUserService interface {
    FetchUsersService() ([]User, error)
}

type externalUserService struct {
    Client JSONPlaceholderClientInterface
}

func NewExternalUserService(client JSONPlaceholderClientInterface) ExternalUserService {
    return &externalUserService{Client: client}
}

func (s *externalUserService) FetchUsersService() ([]User, error) {
    return s.Client.FetchUsers()
}
