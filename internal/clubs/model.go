package clubs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClubID = primitive.ObjectID

type Club struct {
	ID   ClubID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Img  string `json:"img" bson:"img"`
}
