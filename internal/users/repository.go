package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	coll *mongo.Collection
}

func NewMongoUserRepository(c *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{coll: c}
}

func (r *MongoUserRepository) SearchUsers(username string) ([]User, error) {
	filter := bson.M{"username": bson.M{"$regex": username, "$options": "i"}}

	cursor, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return []User{}, err
	}

	users := []User{}
	if err = cursor.All(context.TODO(), &users); err != nil {
		return []User{}, err
	}

	return users, nil
}
