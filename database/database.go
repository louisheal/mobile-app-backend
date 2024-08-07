package database

import (
	"mobile-app-backend/dao"
)

type Database interface {
	GetClubs() ([]dao.Club, error)
}
