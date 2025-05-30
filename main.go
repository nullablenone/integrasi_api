package main

import (
	"context"
	"integrasi_api/config"
	"integrasi_api/internal/domain/user"
	"integrasi_api/internal/integration/jsonplaceholder"
	"integrasi_api/routes"
	"log"
)

func main() {

	// Config
	config.LoadENV()
	db := config.ConnectDB()
	redis := config.ConnectRedis()
	ctx := context.Background()

	err := db.AutoMigrate(user.User{}, user.Address{}, user.Company{})
	if err != nil {
		log.Fatal(err)
	}

	// Api External / Integrasi api
	apiClient := jsonplaceholder.NewJSONPlaceholderClient(config.Env.ExternalAPIURL)
	exUserService := jsonplaceholder.NewExternalUserService(apiClient)

	// Handler User
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(exUserService, userRepository, redis, ctx)
	userHandler := user.NewUserHandler(userService)

	router := routes.SetupRoutes(userHandler)

	router.Run(":808")

}
