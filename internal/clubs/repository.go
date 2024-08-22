package clubs

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClubRepository struct {
	clubs *mongo.Collection
}

func NewMongoClubRepository(db *mongo.Database) *MongoClubRepository {
	return &MongoClubRepository{clubs: db.Collection("clubs")}
}

func (r *MongoClubRepository) GetAllClubs() ([]Club, error) {
	cursor, err := r.clubs.Find(context.TODO(), bson.D{})
	if err != nil {
		return []Club{}, err
	}

	clubs := []Club{}
	err = cursor.All(context.TODO(), &clubs)
	if err != nil {
		return []Club{}, err
	}

	return clubs, nil
}
