package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (routes Routes) GetUsers(c *gin.Context) {
	username := c.Param("username")

	users, err := routes.database.SearchUsers(username)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, users)
}
