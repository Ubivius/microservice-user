package database

import (
	"context"
	"os"
	"testing"

	"github.com/Ubivius/microservice-user/pkg/data"
)

func integrationTestSetup(t *testing.T) {
	t.Log("Test setup")

	if os.Getenv("DB_USERNAME") == "" {
		os.Setenv("DB_USERNAME", "admin")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "pass")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "27888")
	}
	if os.Getenv("DB_HOSTNAME") == "" {
		os.Setenv("DB_HOSTNAME", "localhost")
	}

	err := deleteAllUsersFromMongoDB()
	if err != nil {
		t.Errorf("Failed to delete existing items from database during setup")
	}
}

func addUserAndGetId(t *testing.T) string {
	t.Log("Adding product")
	user := &data.User{
		FirstName:   "testName",
		Username:    "testUsername",
		Email:       "test@email.com",
		DateOfBirth: "01/01/1970",
	}

	mp := NewMongoUsers()
	err := mp.AddUser(context.Background(), user)
	if err != nil {
		t.Errorf("Failed to add product to database")
	}

	t.Log("Fetching new user ID")
	users := mp.GetUsers(context.Background())
	mp.CloseDB()
	return users[len(users)-1].ID
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoUsers()
	if mp == nil {
		t.Error("MongoDB connection is null")
	}
	mp.CloseDB()
}

func TestMongoDBAddUserIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	user := &data.User{
		FirstName:   "testName",
		Username:    "testUsername",
		Email:       "test@email.com",
		DateOfBirth: "01/01/1970",
	}

	mp := NewMongoUsers()
	err := mp.AddUser(context.Background(), user)
	if err != nil {
		t.Errorf("Failed to add User to database")
	}
	users := mp.GetUsers(context.Background())
	if len(users) < 1 {
		t.Errorf("Added user missing from database")
	}
	if len(users) > 1 {
		t.Errorf("User added to database several times during add user call")
	}
	mp.CloseDB()
}

func TestMongoDBUpdateUserIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	userID := addUserAndGetId(t)

	user := &data.User{
		ID:          userID,
		FirstName:   "newName",
		Username:    "testUsername",
		Email:       "test@email.com",
		DateOfBirth: "01/01/1970",
	}

	mp := NewMongoUsers()
	err := mp.UpdateUser(context.Background(), user)
	if err != nil {
		t.Errorf("Error updating user " + err.Error())
	}
	updatedUser, err := mp.GetUserByID(context.Background(), userID)
	if err != nil {
		t.Errorf("Error fetching updated user " + err.Error())
	}
	if updatedUser.FirstName != "newName" {
		t.Errorf("First name updated incorrectly, expected %s but got %s", user.FirstName, updatedUser.FirstName)
	}
	mp.CloseDB()
}

func TestMongoDBGetUsersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	addUserAndGetId(t)

	mp := NewMongoUsers()
	Users := mp.GetUsers(context.Background())
	if Users == nil || len(Users) != 1 {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetUserByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	userID := addUserAndGetId(t)

	mp := NewMongoUsers()
	user, err := mp.GetUserByID(context.Background(), userID)
	if err != nil {
		t.Error("Failed getting user")
	}
	if user == nil {
		t.Error("Returned user is nil")
	}
	if user != nil && user.ID != userID {
		t.Error("Returned incorrect user")
	}

	mp.CloseDB()
}
