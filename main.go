package main

import (
	"mobile-app-backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	mongoDB := database.NewMongoDB()
	routes := Routes{mongoDB}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/clubs", routes.GetClubs)
	router.POST("/ticket", routes.PostTicket)
	router.GET("/tickets", routes.GetTickets)

	router.Run()
}
