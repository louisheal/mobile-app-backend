package clubs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClubHandler struct {
	service *ClubService
}

func NewClubHandler(s *ClubService) *ClubHandler {
	return &ClubHandler{service: s}
}

func (h *ClubHandler) GetClubs(c *gin.Context) {
	clubs, err := h.service.GetAllClubs()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, clubs)
}
