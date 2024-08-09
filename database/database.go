package database

import (
	"mobile-app-backend/dao"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Database interface {
	GetAllClubs() ([]dao.Club, error)
	InsertRating(dao.Rating) error
	GetRatingsFromClubId(clubId primitive.ObjectID) ([]dao.Rating, error)
}
