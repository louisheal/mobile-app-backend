package api

import (
	"mobile-app-backend/internal/clubs"
	"mobile-app-backend/internal/friends"
	"mobile-app-backend/internal/tickets"
	"mobile-app-backend/internal/users"

	"github.com/gin-gonic/gin"
)

func registerTicketRoutes(router *gin.Engine, ticketHandler *tickets.TicketHandler) {
	ticketRoutes := router.Group("/tickets")
	{
		ticketRoutes.GET("", ticketHandler.GetUsersTickets)
		ticketRoutes.POST("", ticketHandler.PostTicket)
		ticketRoutes.PUT("/:ticketID", ticketHandler.PutTicket)
	}
}

func registerClubRoutes(router *gin.Engine, clubHandler *clubs.ClubHandler) {
	clubRoutes := router.Group("/clubs")
	{
		clubRoutes.GET("", clubHandler.GetClubs)
	}
}

func registerUserRoutes(router *gin.Engine, userHandler *users.UserHandler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userHandler.SearchUsers)
	}
}

func registerFriendRoutes(router *gin.Engine, friendHandler *friends.FriendHandler) {
	friendRoutes := router.Group("/friends")
	{
		friendRoutes.POST("", friendHandler.PostFriendRequest)
		friendRoutes.GET("/:fstUser/:sndUser", friendHandler.GetFriendStatus)
		friendRoutes.DELETE("/:fstUser/:sndUser", friendHandler.DeleteFriend)
	}
}

func RegisterRoutes(router *gin.Engine, ticketHandler *tickets.TicketHandler, clubHandler *clubs.ClubHandler, userHandler *users.UserHandler, friendHandler *friends.FriendHandler) {
	registerClubRoutes(router, clubHandler)
	registerTicketRoutes(router, ticketHandler)
	registerUserRoutes(router, userHandler)
	registerFriendRoutes(router, friendHandler)
}
