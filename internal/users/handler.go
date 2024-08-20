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

func (h *UserHandler) GetUsers(c *gin.Context) {
	username := c.Param("username")

	users, err := h.service.SearchUsers(username)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, users)
}
