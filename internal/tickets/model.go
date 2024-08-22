package tickets

import (
	"mobile-app-backend/internal/clubs"
	"mobile-app-backend/internal/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketStatus string

const (
	Valid   = "valid"
	Used    = "used"
	Invalid = "invalid"
)

type TicketID = primitive.ObjectID

type Ticket struct {
	ID     TicketID     `json:"id" bson:"_id"`
	ClubID clubs.ClubID `json:"clubId" bson:"clubId"`
	UserID users.UserID `json:"userId" bson:"userId"`
	Used   bool         `json:"used" bson:"used"`
}

type TicketInput struct {
	ClubID clubs.ClubID `json:"clubId" bson:"clubId"`
	UserID users.UserID `json:"userId" bson:"userId"`
	Used   bool         `json:"used" bson:"used"`
}
