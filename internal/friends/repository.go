package friends

import (
	"context"
	"mobile-app-backend/internal/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFriendRepository struct {
	users   *mongo.Collection
	friends *mongo.Collection
}

func NewMongoFriendRepository(db *mongo.Database) *MongoFriendRepository {
	u, f := db.Collection("users"), db.Collection("friends")
	return &MongoFriendRepository{users: u, friends: f}
}

func (r *MongoFriendRepository) CreateFriend(friend FriendInput) error {
	_, err := r.friends.InsertOne(context.TODO(), friend)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoFriendRepository) FriendExists(sender users.UserID, recipient users.UserID) (bool, error) {
	var friend FriendInput

	filter := bson.M{"sender": sender, "recipient": recipient}
	err := r.friends.FindOne(context.TODO(), filter).Decode(&friend)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MongoFriendRepository) DeleteFriend(fstUser users.UserID, sndUser users.UserID) error {
	filter := bson.M{"sender": fstUser, "recipient": sndUser}
	_, err := r.friends.DeleteOne(context.TODO(), filter)
	return err
}

func (r *MongoFriendRepository) GetUsersFriends(userID users.UserID) ([]users.User, error) {
	filter := bson.M{"recipient": userID}

	cursor, err := r.friends.Find(context.TODO(), filter)
	if err != nil {
		return []users.User{}, err
	}

	friends := []FriendInput{}
	if err = cursor.All(context.TODO(), &friends); err != nil {
		return []users.User{}, err
	}

	result := []users.User{}
	for _, friendInput := range friends {

		if exists, err := r.FriendExists(userID, friendInput.Sender); err != nil {
			return []users.User{}, err
		} else if exists {
			filter := bson.M{"_id": friendInput.Sender}

			var friend users.User
			err := r.users.FindOne(context.TODO(), filter).Decode(&friend)
			if err != nil {
				return []users.User{}, nil
			}

			result = append(result, friend)
		}
	}

	return result, nil
}

func (r *MongoFriendRepository) GetUsersFriendRequests(userID users.UserID) ([]users.User, error) {
	filter := bson.M{"recipient": userID}

	cursor, err := r.friends.Find(context.TODO(), filter)
	if err != nil {
		return []users.User{}, err
	}

	friends := []FriendInput{}
	if err = cursor.All(context.TODO(), &friends); err != nil {
		return []users.User{}, err
	}

	result := []users.User{}
	for _, friendInput := range friends {

		if exists, err := r.FriendExists(userID, friendInput.Sender); err != nil {
			return []users.User{}, err
		} else if !exists {
			filter := bson.M{"_id": friendInput.Sender}

			var friend users.User
			err := r.users.FindOne(context.TODO(), filter).Decode(&friend)
			if err != nil {
				return []users.User{}, nil
			}

			result = append(result, friend)
		}
	}

	return result, nil
}
