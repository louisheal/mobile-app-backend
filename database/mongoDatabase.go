package database

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func (mongoDB *MongoDB) GetClubs() []Club {

	collection := mongoDB.client.Database("mobile-app").Collection("clubs")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var clubs []Club
	cursor.All(context.TODO(), &clubs)
	return clubs
}

func NewMongoDB() *MongoDB {

	password := os.Getenv("dbPass")

	// TODO: Should be passed in (like constructor)
	var uri strings.Builder
	uri.WriteString("mongodb+srv://dbUser:")
	uri.WriteString(password)
	uri.WriteString("@cluster0.iarktte.mongodb.net/?appName=Cluster0")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri.String()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &MongoDB{client}
}
