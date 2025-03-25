package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDatabase initializes MongoDB connection
func ConnectDatabase() {
	mongoURI := os.Getenv("MONGO_URI") // Ensure this is set

	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in the environment variables")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	// Set timeout for connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping the database to check connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB Ping failed:", err)
	}

	DB = client.Database("fiber-mongo-db") // Replace with your DB name
	fmt.Println("🚀 Connected to MongoDB!")
}
