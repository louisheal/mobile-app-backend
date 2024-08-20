package clubs

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClubRepository struct {
	coll *mongo.Collection
}

func NewMongoClubRepository(c *mongo.Collection) *MongoClubRepository {
	return &MongoClubRepository{coll: c}
}

func (db *MongoClubRepository) GetAllClubs() ([]Club, error) {
	cursor, err := db.coll.Find(context.TODO(), bson.D{})
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
