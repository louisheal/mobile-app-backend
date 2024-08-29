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
	database := db.NewMongoDatabase()

	router := setupRouter()

	ticketRepository := tickets.NewMongoTicketRepository(database)
	ticketService := tickets.NewTicketService(ticketRepository)
	ticketHandler := tickets.NewTicketHandler(ticketService)

	clubRepository := clubs.NewMongoClubRepository(database)
	clubService := clubs.NewClubService(clubRepository)
	clubHandler := clubs.NewClubHandler(clubService)

	userRepository := users.NewMongoUserRepository(database)
	userService := users.NewUserService(userRepository)
	userHandler := users.NewUserHandler(userService)

	friendRepository := friends.NewMongoFriendRepository(database)
	friendService := friends.NewFriendService(friendRepository)
	friendHandler := friends.NewFriendHandler(friendService)

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
