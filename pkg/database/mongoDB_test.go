package database

import (
	"testing"

	"github.com/Ubivius/microservice-user/pkg/data"
	"github.com/google/uuid"
)

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

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

	User := &data.User{
		FirstName:   "testName",
		Username:    "testUsername",
		Email:       "test@email.com",
		DateOfBirth: "01/01/1970",
	}

	mp := NewMongoUsers()
	err := mp.AddUser(User)
	if err != nil {
		t.Errorf("Failed to add User to database")
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
	err := mp.UpdateUser(User)
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
	Users := mp.GetUsers()
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
	_, err := mp.GetUserByID("c9ddfb2f-fc4d-40f3-87c0-f6713024a993")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
