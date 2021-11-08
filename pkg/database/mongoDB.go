package database

import (
	"context"
	"fmt"
	"os"

	"github.com/Ubivius/microservice-user/pkg/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// ErrorEnvVar : Environment variable error
var ErrorEnvVar = fmt.Errorf("missing environment variable")

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
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	opts.Monitor = otelmongo.NewMonitor()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	collection := client.Database("ubivius").Collection("users")

	// Assign client and collection to the MongoUsers struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoUsers) PingDB() error {
	return mp.client.Ping(context.Background(), nil)
}

func (mp *MongoUsers) CloseDB() {
	err := mp.client.Disconnect(context.Background())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoUsers) GetUsers(ctx context.Context) data.Users {
	// Users will hold the array of Users
	var users data.Users

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err, "Error getting Users from database")
	}

	// Iterating through cursor
	for cursor.Next(ctx) {
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
	cursor.Close(ctx)

	return users
}

func (mp *MongoUsers) GetUserByID(ctx context.Context, id string) (*data.User, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.User

	// Find a single matching item from the database
	err := mp.collection.FindOne(ctx, filter).Decode(&result)

	// Parse result into the returned User
	return &result, err
}

func (mp *MongoUsers) UpdateUser(ctx context.Context, User *data.User) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: User.ID}}

	// Update sets the matched Users in the database to User
	update := bson.M{"$set": User}

	// Update a single item in the database with the values in update that match the filter
	updateResult, err := mp.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error(err, "Error updating User.")
	}

	if updateResult.MatchedCount != 1 {
		log.Error(data.ErrorUserNotFound, "No matches found for update")
		return err
	}

	return err
}

func (mp *MongoUsers) AddUser(ctx context.Context, User *data.User) error {
	// Inserting the new User into the database
	insertResult, err := mp.collection.InsertOne(ctx, User)
	if err != nil {
		return err
	}

	log.Info("Inserting User", "Inserted ID", insertResult.InsertedID)
	return nil
}

func (mp *MongoUsers) DeleteUser(ctx context.Context, id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err, "Error deleting User")
	}

	log.Info("Deleted documents in Users collection", "delete_count", result.DeletedCount)
	return nil
}

func deleteAllUsersFromMongoDB() error {
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Failing test")
		return err
	}
	collection := client.Database("ubivius").Collection("users")
	_, err = collection.DeleteMany(context.Background(), bson.D{{}})
	return err
}

func mongodbURI() string {
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	if hostname == "" || port == "" || username == "" || password == "" {
		log.Error(ErrorEnvVar, "Some environment variables are not available for the DB connection. DB_HOSTNAME, DB_PORT, DB_USERNAME, DB_PASSWORD")
		os.Exit(1)
	}

	return "mongodb://" + username + ":" + password + "@" + hostname + ":" + port + "/?authSource=admin"
}
