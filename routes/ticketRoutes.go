package routes

import (
	"mobile-app-backend/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (routes Routes) GetTickets(c *gin.Context) {
	userId, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		panic(err)
	}

	tickets, err := routes.database.GetTickets(userId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, tickets)
}

func (routes Routes) PutTicket(c *gin.Context) {
	ticketId, err := primitive.ObjectIDFromHex(c.Query("ticketId"))
	if err != nil {
		// Not valid object id
		c.JSON(http.StatusOK, "Invalid QR Code")
		return
	}

	validTicket, err := routes.database.UseTicket(ticketId)
	if err != nil {
		// Ticket not in database
		c.JSON(http.StatusOK, "Invalid Ticket")
		return
	}

	if validTicket {
		c.JSON(http.StatusOK, "Valid Ticket")
	} else {
		c.JSON(http.StatusOK, "Ticket Already Used")
	}
}

func (routes Routes) PostTicket(c *gin.Context) {
	var newTicket dao.NewTicket
	if err := c.BindJSON(&newTicket); err != nil {
		panic(err)
	}

	id, err := routes.database.CreateTicket(newTicket)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, id)
}
