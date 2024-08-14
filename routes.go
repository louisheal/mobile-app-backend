package main

import (
	"net/http"

	"mobile-app-backend/dao"
	"mobile-app-backend/database"

	"github.com/gin-gonic/gin"
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
