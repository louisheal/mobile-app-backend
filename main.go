package main

import (
	"net/http"

	"mobile-app-backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	database database.Database
}

func main() {

	mongoDB := database.NewMongoDB()
	server := Server{mongoDB}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}

	router.Use(cors.New(config))

	router.GET("/clubs", server.getClubs)

	router.Run()
}

func (server Server) getClubs(c *gin.Context) {
	clubs := server.database.GetClubs()
	c.IndentedJSON(http.StatusOK, clubs)
}
