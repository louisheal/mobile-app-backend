package main

import (
	"net/http"

	"mobile-app-backend/dao"
	"mobile-app-backend/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Routes struct {
	database database.Database
}

func (routes Routes) GetClubs(c *gin.Context) {
	clubs, err := routes.database.GetAllClubs()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, clubs)
}

func (routes Routes) GetTickets(c *gin.Context) {
	tickets, err := routes.database.GetAllTickets()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, tickets)
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

func (routes Routes) PutTicket(c *gin.Context) {
	var ticketId primitive.ObjectID
	if err := c.BindJSON(&ticketId); err != nil {
		panic(err)
	}

	result, err := routes.database.UseTicket(ticketId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, result)
}
