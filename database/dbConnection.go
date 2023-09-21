package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DBInstance() *mongo.Client {
	MongoDB := "mongodb://db:27017/?connect=direct"

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to mongodb")

	return client
}

var Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("restaurant").Collection(collectionName)
}
