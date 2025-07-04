package user

import (
	"context"
	"encoding/json"
	"fmt"
	"integrasi_api/constants"
	"integrasi_api/internal/integration/jsonplaceholder"
	"integrasi_api/internal/kafka/producer"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	ServiceSyncUsers() error
	ServiceGetAllUsers() ([]User, error)
	ProducerUsersSend() error
}

type service struct {
	ExternalUserService jsonplaceholder.ExternalUserService
	Repo                Repository
	Redis               *redis.Client
	Ctx                 context.Context
	Kafka               producer.KafkaProducerInterface
}

func NewUserService(exUserService jsonplaceholder.ExternalUserService, repo Repository, redis *redis.Client, ctx context.Context, kafka producer.KafkaProducerInterface) Service {
	return &service{
		ExternalUserService: exUserService,
		Repo:                repo,
		Redis:               redis,
		Ctx:                 ctx,
		Kafka:               kafka,
	}
}

func (s *service) ServiceSyncUsers() error {
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

func (s *service) ServiceGetAllUsers() ([]User, error) {

	if s.Redis != nil {
		cached, err := s.Redis.Get(s.Ctx, constants.AllUsersCacheKey).Result()
		if err == nil {
			var users []User
			if err := json.Unmarshal([]byte(cached), &users); err == nil {
				return users, nil
			}
		}
	}

	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data user: %w", err)
	}

	if s.Redis != nil {
		data, _ := json.Marshal(users)
		s.Redis.Set(s.Ctx, constants.AllUsersCacheKey, data, 10*time.Minute)
	}

	return users, nil

}

func (s *service) ProducerUsersSend() error {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return fmt.Errorf("gagal mengambil data user: %w", err)
	}

	userData, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("gagal meng-encode data user ke JSON: %w", err)
	}

	// Kirim data JSON ke Kafka
	err = s.Kafka.SendService("users", userData)
	if err != nil {
		return fmt.Errorf("gagal mengirim pesan ke Kafka: %w", err)
	}

	return nil
}
