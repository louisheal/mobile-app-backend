package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"mobile-app-backend/database"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	database database.Database
}

func main() {
	client := connectToMongo()
	mongoDB := database.NewMongoDB(client)
	server := Server{mongoDB}

	router := gin.Default()
	router.GET("/clubs", server.getClubs)
	router.Run()
}

// TODO: Find more appropriate place for this
func connectToMongo() *mongo.Client {
	pass := os.Getenv("dbPass")
	// TODO: String should be in .env (not raw in code)
	uri := fmt.Sprintf("mongodb+srv://dbUser:%s@cluster0.iarktte.mongodb.net/?appName=Cluster0", pass)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func (server Server) getClubs(c *gin.Context) {
	clubs, err := server.database.GetClubs()
	// TODO: Return error one level up or log (need to figure out best practice)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, clubs)
}
