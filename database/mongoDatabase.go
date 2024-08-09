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

const set string = "$set"
const mobileApp string = "mobile-app"
const clubs string = "clubs"
const ratings string = "ratings"

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB(client *mongo.Client) *MongoDB {
	return &MongoDB{client}
}

func (mongoDB *MongoDB) GetAllClubs() ([]dao.Club, error) {
	collection := mongoDB.client.Database(mobileApp).Collection(clubs)

	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "ratings"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "clubId"},
			{Key: "as", Value: "ratings"},
		}}},
		{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$ratings"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$_id"},
			{Key: "name", Value: bson.D{
				{Key: "$first", Value: "$name"},
			}},
			{Key: "rating", Value: bson.D{
				{Key: "$avg", Value: "$ratings.value"},
			}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "name", Value: 1},
			{Key: "rating", Value: bson.D{
				{Key: "$ifNull", Value: bson.A{"$rating", 0}},
			}},
		}}},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return []dao.Club{}, err
	}

	var clubs []dao.Club
	if err = cursor.All(context.TODO(), &clubs); err != nil {
		return nil, fmt.Errorf("failed to decode clubs: %v", err)
	}

	return clubs, nil
}

func (mongoDB *MongoDB) InsertRating(rating dao.Rating) error {
	collection := mongoDB.client.Database(mobileApp).Collection(ratings)

	filter, _ := bson.Marshal(rating.Filter())
	update, _ := bson.Marshal(bson.M{set: rating.Update()})
	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)

	return err
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
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}
