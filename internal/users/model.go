package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserID = primitive.ObjectID

type User struct {
	ID       UserID `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
}
