package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type Club struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Rating float32            `json:"rating" bson:"rating"`
}
