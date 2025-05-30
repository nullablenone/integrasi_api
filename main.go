package main

import (
	"context"
	"integrasi_api/config"
	"integrasi_api/internal/domain/user"
	"integrasi_api/internal/integration/jsonplaceholder"
	"integrasi_api/internal/kafka/producer"
	"integrasi_api/routes"
	"log"
)

func main() {

	// Config
	config.LoadENV()
	ctx := context.Background()
	db := config.ConnectDB()
	redis := config.ConnectRedis()
	config.ConnectBroker("localhost:9092")

	kafkaWriter := config.InitKafkaWriter("localhost:9092", "users")
	defer kafkaWriter.Close()

	err := db.AutoMigrate(user.User{}, user.Address{}, user.Company{})
	if err != nil {
		log.Fatal(err)
	}

	// Api External / Integrasi api
	apiClient := jsonplaceholder.NewJSONPlaceholderClient(config.Env.ExternalAPIURL)
	exUserService := jsonplaceholder.NewExternalUserService(apiClient)

	// Kafka Producer
	producer := producer.NewProducerService(kafkaWriter, ctx)

	// Handler User
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(exUserService, userRepository, redis, ctx, producer)
	userHandler := user.NewUserHandler(userService)

	router := routes.SetupRoutes(userHandler)

	router.Run(":808")

}
