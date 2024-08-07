package database

import (
	"context"
	"fmt"
	"os"

	"mobile-app-backend/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func (mongoDB *MongoDB) GetClubs() ([]dao.Club, error) {
	collection := mongoDB.client.Database("mobile-app").Collection("clubs")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []dao.Club{}, err
	}

	var clubs []dao.Club
	cursor.All(context.TODO(), &clubs)
	return clubs, nil
}

func NewMongoDB(client *mongo.Client) *MongoDB {
	return &MongoDB{client}
}

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
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
