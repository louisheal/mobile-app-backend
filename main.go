package main

import (
	"fmt"
	"net/http"
	"os"

	"mobile-app-backend/database"

	"github.com/gin-gonic/gin"
)

type Server struct {
	database database.Database
}

func main() {

	pass := os.Getenv("dbPass")
	uri := fmt.Sprintf("mongodb+srv://dbUser:%s@cluster0.iarktte.mongodb.net/?appName=Cluster0", pass)

	mongoDB := database.NewMongoDB(uri)
	server := Server{mongoDB}

	router := gin.Default()

	router.GET("/clubs", server.getClubs)
	router.Run()
}

func (server Server) getClubs(c *gin.Context) {
	clubs := server.database.GetClubs()
	c.IndentedJSON(http.StatusOK, clubs)
}
