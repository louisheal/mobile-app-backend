package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	users *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{users: db.Collection("users")}
}

func (r *MongoUserRepository) SearchUsers(username string) ([]User, error) {
	filter := bson.M{"username": bson.M{"$regex": username, "$options": "i"}}

	cursor, err := r.users.Find(context.TODO(), filter)
	if err != nil {
		return []User{}, err
	}

	users := []User{}
	if err = cursor.All(context.TODO(), &users); err != nil {
		return []User{}, err
	}

	return users, nil
}
