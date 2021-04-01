package database

import (
	"context"
	"os"

	"github.com/Ubivius/microservice-user/pkg/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUsers struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUsers() UserDB {
	mp := &MongoUsers{}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		log.Error(err, "MongoDB setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *MongoUsers) Connect() error {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:pass@localhost:27888/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	collection := client.Database("ubivius").Collection("Users")

	// Assign client and collection to the MongoUsers struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoUsers) CloseDB() {
	err := mp.client.Disconnect(context.TODO())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoUsers) GetUsers() data.Users {
	// Users will hold the array of Users
	var users data.Users

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error(err, "Error getting Users from database")
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var result data.User
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding User from database")
		}
		users = append(users, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return users
}

func (mp *MongoUsers) GetUserByID(id string) (*data.User, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.User

	// Find a single matching item from the database
	err := mp.collection.FindOne(context.TODO(), filter).Decode(&result)

	// Parse result into the returned User
	return &result, err
}

func (mp *MongoUsers) UpdateUser(User *data.User) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: User.ID}}

	// Update sets the matched Users in the database to User
	update := bson.M{"$set": User}

	// Update a single item in the database with the values in update that match the filter
	_, err := mp.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error(err, "Error updating User.")
	}

	return err
}

func (mp *MongoUsers) AddUser(User *data.User) error {
	// Inserting the new User into the database
	insertResult, err := mp.collection.InsertOne(context.TODO(), User)
	if err != nil {
		return err
	}

	log.Info("Inserting User", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoUsers) DeleteUser(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error deleting User")
	}

	log.Info("Deleted documents in Users collection", "delete_count", result.DeletedCount)
	return nil
}
