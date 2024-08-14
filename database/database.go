package database

import (
	"mobile-app-backend/dao"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Database interface {
	GetAllClubs() ([]dao.Club, error)
	CreateTicket(dao.NewTicket) (primitive.ObjectID, error)
}
