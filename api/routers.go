package api

import (
	"mobile-app-backend/internal/clubs"
	"mobile-app-backend/internal/friends"
	"mobile-app-backend/internal/tickets"
	"mobile-app-backend/internal/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, ticketHandler *tickets.TicketHandler, clubHandler *clubs.ClubHandler,
	userHandler *users.UserHandler, friendHandler *friends.FriendHandler) {
	ticketRoutes := router.Group("/tickets")
	{
		ticketRoutes.GET("", ticketHandler.GetTickets)
		ticketRoutes.POST("", ticketHandler.PostTicket)
		ticketRoutes.PUT("/:ticketID", ticketHandler.PutTicket)
	}
	clubRoutes := router.Group("/clubs")
	{
		clubRoutes.GET("", clubHandler.GetClubs)
	}
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userHandler.GetUsers)
	}
	friendRoutes := router.Group("/friends")
	{
		friendRoutes.POST("", friendHandler.PostFriendRequest)
		friendRoutes.GET("/:fstUser/:sndUser", friendHandler.GetFriendStatus)
	}
}
