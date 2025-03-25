package repositories

import (
	"context"
	"go-fiber-mongo-crud/config"
	"go-fiber-mongo-crud/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

// Ensure MongoDB is initialized before using the collection
func getUserCollection() *mongo.Collection {
	if config.DB == nil {
		log.Fatal("Database is not initialized. Ensure ConnectDatabase() is called first.")
	}
	if userCollection == nil {
		userCollection = config.DB.Collection("users")
	}
	return userCollection
}

// Create a new user
// Create a new user and return the created user
func CreateUser(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getUserCollection()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// Set the generated ID back to the user object
	user.ID = result.InsertedID.(primitive.ObjectID)

	return &user, nil
}

// Get all users
func GetUsers() ([]models.User, error) {
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := getUserCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID fetches a user by their ID
func GetUserByID(userID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getUserCollection()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user by their ID
func UpdateUser(userID string, updatedData models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getUserCollection()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedData}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	// Fetch updated user
	return GetUserByID(userID)
}

// DeleteUser removes a user by their ID
func DeleteUser(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := getUserCollection()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
