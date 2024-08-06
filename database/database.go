package database

type Database interface {
	GetClubs() ([]Club, error)
}
