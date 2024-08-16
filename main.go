package main

import (
	"mobile-app-backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := database.NewMongoDB()
	routes := Routes{database}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/clubs", routes.GetClubs)

	router.GET("/tickets/:userId", routes.GetTickets)
	router.POST("/ticket", routes.PostTicket)
	router.PUT("/ticket", routes.PutTicket)

	router.Run()
}
