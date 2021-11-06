package database

import (
	"context"
	"os"
	"testing"

	"github.com/Ubivius/microservice-user/pkg/data"
	"github.com/google/uuid"
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

func addProductAndGetId(t *testing.T) string {
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
	products := mp.GetUsers(context.Background())
	mp.CloseDB()
	return products[len(products)-1].ID
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoUsers()
	if mp == nil {
		t.Fail()
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

	User := &data.User{
		ID:          uuid.NewString(),
		FirstName:   "testName",
		Username:    "testUsername",
		Email:       "test@email.com",
		DateOfBirth: "01/01/1970",
	}

	mp := NewMongoUsers()
	err := mp.UpdateUser(context.Background(), User)
	if err != nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBGetUsersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoUsers()
	Users := mp.GetUsers(context.Background())
	if Users == nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetUserByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoUsers()
	_, err := mp.GetUserByID(context.Background(), "c9ddfb2f-fc4d-40f3-87c0-f6713024a993")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
