package routes

import (
	"integrasi_api/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler *user.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/get-users", userHandler.GetAllUsers)

	return router

}
