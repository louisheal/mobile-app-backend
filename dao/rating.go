package dao

type Rating struct {
	ClubID string `json:"clubId" bson:"clubId"`
	UserID string `json:"userId" bson:"userId"`
	Value  int    `json:"value" bson:"value"`
}
