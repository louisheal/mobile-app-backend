package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type Club struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Img    string             `json:"img" bson:"img"`
	Rating float32            `json:"rating" bson:"rating"`
}
