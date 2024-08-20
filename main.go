package main

import (
	"mobile-app-backend/database"
	"mobile-app-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := database.NewMongoDB()
	routes := routes.NewRoutes(database)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/clubs", routes.GetClubs)

	router.GET("/tickets/:userId", routes.GetTickets)
	router.POST("/ticket", routes.PostTicket)
	router.PUT("/ticket", routes.PutTicket)

	router.GET("/users/:username", routes.GetUsers)

	router.POST("/friends", routes.PostFriendRequest)
	router.GET("/friends/:fstUser/:sndUser", routes.GetFriendRequestStatus)

	router.Run()
}
