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

	c.IndentedJSON(http.StatusOK, clubs)
}

func (routes Routes) PutRating(c *gin.Context) {
	var newRating dao.Rating
	if err := c.BindJSON(&newRating); err != nil {
		panic(err)
	}

	if err := routes.database.InsertRating(newRating); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, newRating)
}
