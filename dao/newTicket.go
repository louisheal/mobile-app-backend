package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewTicket struct {
	ClubID primitive.ObjectID `json:"clubId" bson:"clubId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	Used   bool               `json:"used" bson:"used"`
}
