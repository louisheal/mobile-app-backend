package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	username := c.Query("username")

	users, err := h.service.SearchUsers(username)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, users)
}
