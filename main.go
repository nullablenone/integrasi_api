package main

import (
	"fmt"
	"integrasi_api/config"
	"integrasi_api/internal/domain/user"
	"integrasi_api/internal/integration/jsonplaceholder"
	"integrasi_api/routes"
)

func main() {
	config.LoadENV()
	config.ConnectDB()

	apiClient := jsonplaceholder.NewJSONPlaceholderClient(config.Env.ExternalAPIURL)
	exUserService := jsonplaceholder.NewExternalUserService(apiClient)

	userService := user.NewUserService(exUserService)
	userHandler := user.NewUserHandler(userService)

	router := routes.SetupRoutes(userHandler)


	router.Run(":8080")

	fmt.Println()
}
