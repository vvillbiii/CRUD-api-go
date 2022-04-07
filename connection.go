package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connection() *mongo.Database{

	//checking if .env file exist
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")


	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}


	fmt.Println("Successfully connected and pinged.")

	//specifying database
	db := client.Database("go_watch_CRUD_app")

	return db
}