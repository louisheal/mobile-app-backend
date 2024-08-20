package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketHandler struct {
	service *TicketService
}

func NewTicketHandler(s *TicketService) *TicketHandler {
	return &TicketHandler{service: s}
}

func (h *TicketHandler) GetTickets(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Query("userID"))
	if err != nil {
		panic(err)
	}

	tickets, err := h.service.GetUsersTickets(userID)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) PostTicket(c *gin.Context) {
	var newTicket NewTicket
	if err := c.BindJSON(&newTicket); err != nil {
		panic(err)
	}

	id, err := h.service.CreateTicket(newTicket)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, id)
}

func (h *TicketHandler) PutTicket(c *gin.Context) {
	ticketId, err := primitive.ObjectIDFromHex(c.Query("ticketId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	status := h.service.UseTicket(ticketId)
	c.JSON(http.StatusOK, status)
}
