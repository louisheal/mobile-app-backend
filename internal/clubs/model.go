package clubs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Club struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Img  string             `json:"img" bson:"img"`
}

// type ClubID primitive.ObjectID

// func (id ClubID) MarshalJSON() ([]byte, error) {
// 	oid := primitive.ObjectID(id)
// 	return oid.MarshalJSON()
// }

// func (id *ClubID) UnmarshalJSON(data []byte) error {
// 	var oid primitive.ObjectID
// 	err := oid.UnmarshalJSON(data)
// 	if err != nil {
// 		return err
// 	}

// 	*id = ClubID(oid)

// 	return nil
// }

// func (id ClubID) MarshalBSON() ([]byte, error) {
// 	oid := primitive.ObjectID(id)
// 	return oid.MarshalJSON()
// }

// func (id *ClubID) UnmarshalBSON(data []byte) error {
// 	var oid primitive.ObjectID
// 	err := oid.UnmarshalJSON(data)
// 	if err != nil {
// 		return err
// 	}

// 	*id = ClubID(oid)

// 	return nil
// }
