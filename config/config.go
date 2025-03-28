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
	mongoURI := os.Getenv("MONGO_URI") // Read from environment variable

	if mongoURI == "" {
		log.Fatal("❌ MONGO_URI is not set in the environment variables")
	}

	// Add TLS configuration explicitly
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(10 * time.Second)

	// Create connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ Error connecting to MongoDB: %v", err)
	}

	// Ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("❌ MongoDB Ping failed: %v", err)
	}

	DB = client.Database("fiber-mongo-db") // ✅ Make sure this DB name exists
	fmt.Println("🚀 Connected to MongoDB successfully!")
}
