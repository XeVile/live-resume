package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance ... Function to create a database instance
func DBinstance() *mongo.Client {
  // Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

  // Collect system registry for MONGODB_URL
	MongoDB := os.Getenv("MONGODB_URL")

  // Create new client with applied URL options
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

  // Set timeout for a blocking function request 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

  // Call cancel only after DB instantiation
	defer cancel()

  // Attempt connection request
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// Client ... New Database instance
var Client *mongo.Client = DBinstance()

// OpenCollection ... is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)

	return collection
}
