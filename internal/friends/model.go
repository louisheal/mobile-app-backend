package friends

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendInput struct {
	Sender    primitive.ObjectID `json:"sender" bson:"sender"`
	Recipient primitive.ObjectID `json:"recipient" bson:"recipient"`
}

type FriendStatus string

const (
	Accepted = "accepted"
	Accept   = "accept"
	Pending  = "pending"
	Send     = "send"
	None     = ""
)
