package database

import (
	"context"
	"fmt"
	"os"

	"mobile-app-backend/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mobileApp string = "mobile-app"
const clubs string = "clubs"
const tickets string = "tickets"
const users string = "users"

type MongoDB struct {
	client *mongo.Client
}

func (mongoDB *MongoDB) GetAllClubs() ([]dao.Club, error) {
	collection := mongoDB.client.Database(mobileApp).Collection(clubs)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []dao.Club{}, err
	}

	clubs := []dao.Club{}
	if err = cursor.All(context.TODO(), &clubs); err != nil {
		return []dao.Club{}, err
	}

	return clubs, nil
}

func (mongoDB *MongoDB) GetTickets(userId primitive.ObjectID) ([]dao.Ticket, error) {
	collection := mongoDB.client.Database(mobileApp).Collection(tickets)

	cursor, err := collection.Find(context.TODO(), bson.M{"userId": userId})
	if err != nil {
		return []dao.Ticket{}, err
	}

	tickets := []dao.Ticket{}
	if err = cursor.All(context.TODO(), &tickets); err != nil {
		return []dao.Ticket{}, err
	}

	return tickets, nil
}

func (mongoDB *MongoDB) CreateTicket(newTicket dao.NewTicket) (primitive.ObjectID, error) {
	collection := mongoDB.client.Database(mobileApp).Collection(tickets)

	result, err := collection.InsertOne(context.TODO(), newTicket)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return id, err
}

func (mongoDB *MongoDB) UseTicket(ticketId primitive.ObjectID) (bool, error) {
	collection := mongoDB.client.Database(mobileApp).Collection(tickets)

	filter := bson.M{"_id": ticketId}

	var ticket dao.Ticket
	err := collection.FindOne(context.TODO(), filter).Decode(&ticket)
	if err != nil {
		return false, err
	}
	if ticket.Used {
		return false, nil
	}

	update := bson.M{"$set": bson.M{"used": true}}
	opts := options.Update().SetUpsert(true)

	if _, err = collection.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return false, err
	}

	return true, nil
}

func (mongoDB *MongoDB) SearchUsers(username string) ([]dao.User, error) {
	collection := mongoDB.client.Database(mobileApp).Collection(users)

	filter := bson.M{"username": bson.M{"$regex": username, "$options": "i"}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []dao.User{}, err
	}

	var users []dao.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return []dao.User{}, err
	}

	return users, nil
}

func NewMongoDB() *MongoDB {
	pass := os.Getenv("dbPass")
	// TODO: String should be in .env (not raw in code)
	uri := fmt.Sprintf("mongodb+srv://dbUser:%s@cluster0.iarktte.mongodb.net/?appName=Cluster0", pass)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	// TODO: Use logging instead
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &MongoDB{client}
}
