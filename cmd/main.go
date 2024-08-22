package main

import (
	"mobile-app-backend/api"
	"mobile-app-backend/db"
	"mobile-app-backend/internal/clubs"
	"mobile-app-backend/internal/friends"
	"mobile-app-backend/internal/tickets"
	"mobile-app-backend/internal/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	database := db.NewMongoDatabase()

	router := setupRouter()

	ticketHandler := setupTicketHandler(database.Collection("tickets"))
	clubHandler := setupClubHandler(database.Collection("clubs"))
	userHandler := setupUserHandler(database.Collection("users"))
	friendHandler := setupFriendHandler(database.Collection("friends"))

	api.RegisterRoutes(router, ticketHandler, clubHandler, userHandler, friendHandler)

	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	router.Use(cors.New(config))

	return router
}

func setupTicketHandler(coll *mongo.Collection) *tickets.TicketHandler {
	ticketRepository := tickets.NewMongoTicketRepository(coll)
	ticketService := tickets.NewTicketService(ticketRepository)
	return tickets.NewTicketHandler(ticketService)
}

func setupClubHandler(coll *mongo.Collection) *clubs.ClubHandler {
	clubRepository := clubs.NewMongoClubRepository(coll)
	clubService := clubs.NewClubService(clubRepository)
	return clubs.NewClubHandler(clubService)
}

func setupUserHandler(coll *mongo.Collection) *users.UserHandler {
	userRepository := users.NewMongoUserRepository(coll)
	userService := users.NewUserService(userRepository)
	return users.NewUserHandler(userService)
}

func setupFriendHandler(coll *mongo.Collection) *friends.FriendHandler {
	friendRepository := friends.NewMongoFriendRepository(coll)
	friendService := friends.NewFriendService(friendRepository)
	return friends.NewFriendHandler(friendService)
}
