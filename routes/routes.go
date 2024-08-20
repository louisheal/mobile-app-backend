package routes

import (
	"mobile-app-backend/database"
)

type Routes struct {
	database database.Database
}

func NewRoutes(database database.Database) *Routes {
	return &Routes{database}
}
