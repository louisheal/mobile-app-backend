package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var clubsPipeline mongo.Pipeline = mongo.Pipeline{
	{{lookup, bson.D{
		{from, ratings},
		{localField, id},
		{foreignField, clubId},
		{as, ratings},
	}}},
	{{unwind, bson.D{
		{path, "$ratings"},
		{preserveNullAndEmptyArrays, true},
	}}},
	{{Key: group, Value: bson.D{
		{id, "$_id"},
		{name, bson.D{
			{first, "$name"},
		}},
		{img, bson.D{
			{first, "$img"},
		}},
		{rating, bson.D{
			{avg, "$ratings.value"},
		}},
	}}},
	{{project, bson.D{
		{id, 1},
		{name, 1},
		{img, 1},
		{rating, bson.D{
			{ifNull, bson.A{"$rating", 0}},
		}},
	}}},
}
