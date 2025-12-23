package config

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	oc     sync.Once
)

func Connect() *mongo.Client {
	oc.Do(func() {
		mongoURI := GetConfig().GetConnectionURL()
		clientOptions := options.Client().ApplyURI(mongoURI)

		var err error
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		if err = client.Ping(context.Background(), nil); err != nil {
			log.Fatal(err)
		}
	})

	log.Println("Connected to MongoDB")
	return client
}

func GetDatabase() *mongo.Client {
	if client == nil {
		client = Connect()
	}

	return client
}

func Close(client *mongo.Client) {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect MongoDB client: %v", err)
	}
}

func GetUserCollection(client *mongo.Client) *mongo.Collection {
	name := GetConfig().GetDatabaseName()
	return client.Database(name).Collection("users")
}

func GetUrlStoreCollection(client *mongo.Client) *mongo.Collection {
	name := GetConfig().GetDatabaseName()
	return client.Database(name).Collection("url_store")
}
