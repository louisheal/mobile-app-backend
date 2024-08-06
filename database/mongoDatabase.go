package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	client *mongo.Client
}

func (mongoDB *MongoDB) GetClubs() ([]Club, error) {
	collection := mongoDB.client.Database("mobile-app").Collection("clubs")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []Club{}, err
	}

	var clubs []Club
	cursor.All(context.TODO(), &clubs)
	return clubs, nil
}

func NewMongoDB(client *mongo.Client) *MongoDB {
	return &MongoDB{client}
}
