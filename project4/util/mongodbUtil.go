package util

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	URL = "mongodb://localhost:27017"
)

func GetSession() (*mongo.Client, error, *mongo.Collection) {
	clientOptions := options.Client().ApplyURI(URL)
	client, ero := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("mydb").Collection("user")
	return client, ero, collection
}
