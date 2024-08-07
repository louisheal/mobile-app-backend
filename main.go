package main

import (
	"mobile-app-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	client := database.ConnectToMongo()
	mongoDB := database.NewMongoDB(client)
	routes := Routes{mongoDB}

	router := gin.Default()
	router.GET("/clubs", routes.GetClubs)
	router.Run()
}
