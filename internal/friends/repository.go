package friends

import (
	"context"
	"mobile-app-backend/internal/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFriendRepository struct {
	coll *mongo.Collection
}

func NewMongoFriendRepository(c *mongo.Collection) *MongoFriendRepository {
	return &MongoFriendRepository{coll: c}
}

func (r *MongoFriendRepository) CreateFriend(friend FriendInput) error {
	_, err := r.coll.InsertOne(context.TODO(), friend)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoFriendRepository) FriendExists(fstUser users.UserID, sndUser users.UserID) (bool, error) {
	var friend FriendInput

	filter := bson.M{"sender": fstUser, "recipient": sndUser}
	err := r.coll.FindOne(context.TODO(), filter).Decode(&friend)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MongoFriendRepository) DeleteFriend(fstUser users.UserID, sndUser users.UserID) error {
	filter := bson.M{"sender": fstUser, "recipient": sndUser}
	_, err := r.coll.DeleteOne(context.TODO(), filter)
	return err
}
