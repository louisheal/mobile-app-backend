package main

import (
	"context"
	"fmt"
	"mobile-app-backend/api"
	"mobile-app-backend/internal/clubs"
	"mobile-app-backend/internal/friends"
	"mobile-app-backend/internal/tickets"
	"mobile-app-backend/internal/users"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoDatabase := NewMongoDB()

	ticketCollection := mongoDatabase.Database("mobile-app").Collection("tickets")
	ticketRepository := tickets.NewMongoTicketRepository(ticketCollection)
	ticketService := tickets.NewTicketService(ticketRepository)
	ticketHandler := tickets.NewTicketHandler(ticketService)

	clubCollection := mongoDatabase.Database("mobile-app").Collection("clubs")
	clubRespository := clubs.NewMongoClubRepository(clubCollection)
	clubService := clubs.NewClubService(clubRespository)
	clubHandler := clubs.NewClubHandler(clubService)

	userCollection := mongoDatabase.Database("mobile-app").Collection("users")
	userRepository := users.NewMongoUserRepository(userCollection)
	userService := users.NewUserService(userRepository)
	userHandler := users.NewUserHandler(userService)

	friendCollection := mongoDatabase.Database("mobile-app").Collection("friends")
	friendRespository := friends.NewMongoFriendRepository(friendCollection)
	friendService := friends.NewFriendService(friendRespository)
	friendHandler := friends.NewFriendHandler(friendService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	api.RegisterRoutes(router, ticketHandler, clubHandler, userHandler, friendHandler)

	router.Run()
}

func NewMongoDB() *mongo.Client {
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
	// TODO: Use logging instead
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
