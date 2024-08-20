package tickets

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketStatus string

const (
	Valid   = "valid"
	Used    = "used"
	Invalid = "invalid"
)

type Ticket struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	ClubID primitive.ObjectID `json:"clubId" bson:"clubId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	Used   bool               `json:"used" bson:"used"`
}

type NewTicket struct {
	ClubID primitive.ObjectID `json:"clubId" bson:"clubId"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	Used   bool               `json:"used" bson:"used"`
}
