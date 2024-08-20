package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (routes Routes) GetClubs(c *gin.Context) {
	clubs, err := routes.database.GetAllClubs()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, clubs)
}
