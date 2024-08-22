package friends

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *MongoFriendRepository) FriendExists(fstUser primitive.ObjectID, sndUser primitive.ObjectID) (bool, error) {
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
