package main

import (
	"mobile-app-backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	client := database.ConnectToMongo()
	mongoDB := database.NewMongoDB(client)
	routes := Routes{mongoDB}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	router.GET("/clubs", routes.GetClubs)
	router.PUT("/rating", routes.PutRating)

	router.Run()
}
