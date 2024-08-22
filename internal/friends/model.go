package friends

import (
	"mobile-app-backend/internal/users"
)

type FriendInput struct {
	Sender    users.UserID `json:"sender" bson:"sender"`
	Recipient users.UserID `json:"recipient" bson:"recipient"`
}

type FriendStatus string

const (
	Accepted = "accepted"
	Accept   = "accept"
	Pending  = "pending"
	Send     = "send"
	None     = ""
)
