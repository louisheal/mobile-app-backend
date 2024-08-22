package tickets

import (
	"context"
	"mobile-app-backend/internal/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTicketRepository struct {
	tickets *mongo.Collection
}

func NewMongoTicketRepository(db *mongo.Database) *MongoTicketRepository {
	return &MongoTicketRepository{tickets: db.Collection("tickets")}
}

func (r *MongoTicketRepository) GetUsersTickets(userId users.UserID) ([]Ticket, error) {
	cursor, err := r.tickets.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		return []Ticket{}, err
	}

	result := []Ticket{}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return []Ticket{}, err
	}

	return result, nil
}

func (r *MongoTicketRepository) GetTicket(ticketId TicketID) (Ticket, error) {
	filter := bson.M{"_id": ticketId}

	var result Ticket
	err := r.tickets.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return Ticket{}, err
	}

	return result, nil
}

func (r *MongoTicketRepository) CreateTicket(newTicket TicketInput) (TicketID, error) {

	result, err := r.tickets.InsertOne(context.TODO(), newTicket)
	if err != nil {
		return TicketID{}, err
	}

	id := result.InsertedID.(TicketID)

	return id, nil
}

func (r *MongoTicketRepository) UseTicket(ticketId TicketID) error {
	filter := bson.M{"_id": ticketId}
	update := bson.M{"$set": bson.M{"used": true}}
	opts := options.Update().SetUpsert(true)

	if _, err := r.tickets.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return err
	}

	return nil
}
