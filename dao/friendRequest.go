package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendRequest struct {
	Sender    primitive.ObjectID `json:"sender" bson:"sender"`
	Recipient primitive.ObjectID `json:"recipient" bson:"recipient"`
}
