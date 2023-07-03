package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var client *mongo.Client

func connectToMongo() {
	databaseConfiguration := GetDatabaseConfiguration()
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseConfiguration.ConnectionString))
	if err != nil {
		log.Fatal("Could not able to connect to the database, Reason:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Context error, mongoDB:", err)
	}

	//Cancel context to avoid memory leak
	defer cancel()

	DB = client.Database(databaseConfiguration.DatabaseName)

	return
}

func disconnectFromMongo() {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

// In Golang, init() functions always initialize whenever the package is called.
// So, whenever DB variable called, the init() function initialized
func init() {
	connectToMongo()
}
