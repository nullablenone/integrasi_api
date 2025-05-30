package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService Service
}

func NewUserHandler(userService Service) *Handler {
	return &Handler{UserService: userService}
}

func (h *Handler) SyncUsers(c *gin.Context) {
	// manggil service
	users, err := h.UserService.SyncUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "err bro",
		})
	}

	for _, u := range users {
		fmt.Println(u.Name)
	}

}
