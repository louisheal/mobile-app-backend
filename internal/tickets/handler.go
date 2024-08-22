package tickets

import (
	"fmt"
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

func (h *TicketHandler) GetUsersTickets(c *gin.Context) {
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
	var ticket TicketInput
	if err := c.BindJSON(&ticket); err != nil {
		panic(err)
	}

	fmt.Println(ticket.ClubID)

	id, err := h.service.CreateTicket(ticket)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

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
