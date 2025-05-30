package user

import (
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
	if err := h.UserService.ServiceSyncUsers(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "err bro",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "berhasil",
	})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.UserService.ServiceGetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "err bro",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}
