package database

import (
	"mobile-app-backend/dao"
)

type Database interface {
	GetAllClubs() ([]dao.Club, error)
	InsertRating(dao.Rating) error
}
