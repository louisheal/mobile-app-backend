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

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB(client *mongo.Client) *MongoDB {
	return &MongoDB{client}
}

func (mongoDB *MongoDB) GetAllClubs() ([]dao.Club, error) {
	collection := mongoDB.client.Database(mobileApp).Collection(clubs)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []dao.Club{}, err
	}

	var clubs []dao.Club
	if err = cursor.All(context.TODO(), &clubs); err != nil {
		return []dao.Club{}, err
	}

	return clubs, nil
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

// TODO: Function feels like it should be in another file
func ConnectToMongo() *mongo.Client {
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

	return client
}
