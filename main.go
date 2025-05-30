package main

import (
	"fmt"
	"integrasi_api/config"
	"integrasi_api/internal/domain/user"
	"integrasi_api/internal/integration/jsonplaceholder"
	"integrasi_api/routes"
	"log"
)

func main() {
	config.LoadENV()
	db := config.ConnectDB()

	err := db.AutoMigrate(user.User{}, user.Address{}, user.Company{})
	if err != nil {
		log.Fatal(err)
	}

	apiClient := jsonplaceholder.NewJSONPlaceholderClient(config.Env.ExternalAPIURL)
	exUserService := jsonplaceholder.NewExternalUserService(apiClient)

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(exUserService, userRepository)
	userHandler := user.NewUserHandler(userService)

	router := routes.SetupRoutes(userHandler)

	router.Run(":808")

	fmt.Println()
}
