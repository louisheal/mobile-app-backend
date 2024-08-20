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
)

func main() {
	mongoDatabase := db.NewMongoClient().Database("mobile-app")

	ticketCollection := mongoDatabase.Collection("tickets")
	ticketRepository := tickets.NewMongoTicketRepository(ticketCollection)
	ticketService := tickets.NewTicketService(ticketRepository)
	ticketHandler := tickets.NewTicketHandler(ticketService)

	clubCollection := mongoDatabase.Collection("clubs")
	clubRespository := clubs.NewMongoClubRepository(clubCollection)
	clubService := clubs.NewClubService(clubRespository)
	clubHandler := clubs.NewClubHandler(clubService)

	userCollection := mongoDatabase.Collection("users")
	userRepository := users.NewMongoUserRepository(userCollection)
	userService := users.NewUserService(userRepository)
	userHandler := users.NewUserHandler(userService)

	friendCollection := mongoDatabase.Collection("friends")
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
