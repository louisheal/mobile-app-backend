package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rating struct {
	ClubID primitive.ObjectID `json:"clubId" bson:"clubId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	Value  int                `json:"value" bson:"value"`
}

type RatingFilter struct {
	ClubID primitive.ObjectID `json:"clubId" bson:"clubId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
}

type RatingUpdate struct {
	Value int `json:"value" bson:"value"`
}

func (rating Rating) Filter() RatingFilter {
	return RatingFilter{rating.ClubID, rating.UserID}
}

func (rating Rating) Update() RatingUpdate {
	return RatingUpdate{rating.Value}
}
