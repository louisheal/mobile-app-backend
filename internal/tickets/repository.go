package tickets

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTicketRepository struct {
	coll *mongo.Collection
}

func NewMongoTicketRepository(c *mongo.Collection) *MongoTicketRepository {
	return &MongoTicketRepository{coll: c}
}

func (r *MongoTicketRepository) GetUsersTickets(userId primitive.ObjectID) ([]Ticket, error) {
	cursor, err := r.coll.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		return []Ticket{}, err
	}

	result := []Ticket{}
	if err = cursor.All(context.TODO(), &result); err != nil {
		return []Ticket{}, err
	}

	return result, nil
}

func (r *MongoTicketRepository) GetTicket(ticketId primitive.ObjectID) (Ticket, error) {
	filter := bson.M{"_id": ticketId}

	var result Ticket
	err := r.coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return Ticket{}, err
	}

	return result, nil
}

func (r *MongoTicketRepository) CreateTicket(newTicket NewTicket) (primitive.ObjectID, error) {
	result, err := r.coll.InsertOne(context.TODO(), newTicket)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return id, nil
}

func (r *MongoTicketRepository) UseTicket(ticketId primitive.ObjectID) error {
	filter := bson.M{"_id": ticketId}
	update := bson.M{"$set": bson.M{"used": true}}
	opts := options.Update().SetUpsert(true)

	if _, err := r.coll.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return err
	}

	return nil
}
