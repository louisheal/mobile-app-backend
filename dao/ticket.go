package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	ClubID primitive.ObjectID `json:"clubId" bson:"clubId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
}
