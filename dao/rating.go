package dao

type Rating struct {
	ClubID string `json:"clubId" bson:"clubId"`
	UserID string `json:"userId" bson:"userId"`
	Value  int    `json:"value" bson:"value"`
}

type RatingFilter struct {
	ClubID string `json:"clubId" bson:"clubId"`
	UserID string `json:"userId" bson:"userId"`
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
