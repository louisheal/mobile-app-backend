package database

import (
	"mobile-app-backend/dao"
	"mobile-app-backend/friendStatus"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Database interface {
	GetAllClubs() ([]dao.Club, error)

	GetTickets(primitive.ObjectID) ([]dao.Ticket, error)
	CreateTicket(dao.NewTicket) (primitive.ObjectID, error)
	UseTicket(primitive.ObjectID) (bool, error)

	SearchUsers(string) ([]dao.User, error)

	SendFriendRequest(dao.FriendRequest) error
	GetFriendRequestStatus(primitive.ObjectID, primitive.ObjectID) (friendStatus.FriendStatus, error)
}
