package main

import (
	"net/http"

	"mobile-app-backend/database"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	database database.Database
}

func (routes Routes) GetClubs(c *gin.Context) {
	clubs, err := routes.database.GetClubs()
	// TODO: Return error one level up or log (need to figure out best practice)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, clubs)
}
