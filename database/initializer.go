package database

import (
	"context"
	"time"

	"jwt-project/common/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *mongo.Client

func Connect() *mongo.Client {
	if instance != nil {
		return instance
	}

	mongoClient, _ := mongo.NewClient(options.Client().ApplyURI(env.MONGO_URL))

	timeout := 10 * time.Second
	connection, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	mongoClient.Connect(connection)
	return mongoClient
}

func Collection(client *mongo.Client, collectionName string) *mongo.Collection {
	var createCollection *mongo.Collection = client.Database(env.MONGO_COLLECTION_NAME).Collection(collectionName)

	return createCollection
}
