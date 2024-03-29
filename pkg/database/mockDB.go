package database

import (
	"context"

	"github.com/Ubivius/microservice-user/pkg/data"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

type MockUsers struct {
}

func NewMockUsers() UserDB {
	log.Info("Connecting to mock database")
	return &MockUsers{}
}

func (mp *MockUsers) Connect() error {
	return nil
}

func (mp *MockUsers) PingDB() error {
	return nil
}

func (mp *MockUsers) CloseDB() {
	log.Info("Mocked DB connection closed")
}

//Return every users
func (mp *MockUsers) GetUsers(ctx context.Context) data.Users {
	_, span := otel.Tracer("users").Start(ctx, "getUsersDatabase")
	defer span.End()
	return userList
}

// GetUserByID returns a single user with the given id
func (mp *MockUsers) GetUserByID(ctx context.Context, id string) (*data.User, error) {
	_, span := otel.Tracer("users").Start(ctx, "getUserByIdDatabase")
	defer span.End()
	index := findIndexByUserID(id)
	if index == -1 {
		return nil, data.ErrorUserNotFound
	}
	return userList[index], nil
}

// GetUserByUsername returns a single user with the given id
func (mp *MockUsers) GetUserByUsername(ctx context.Context, username string) (*data.User, error) {
	_, span := otel.Tracer("users").Start(ctx, "getUserByUsernameDatabase")
	defer span.End()
	index := findIndexByUsername(username)
	if index == -1 {
		return nil, data.ErrorUserNotFound
	}
	return userList[index], nil
}

// AddUser creates a new user
func (mp *MockUsers) AddUser(ctx context.Context, user *data.User) error {
	_, span := otel.Tracer("users").Start(ctx, "addUserDatabase")
	defer span.End()
	user.ID = uuid.NewString()
	userList = append(userList, user)
	return nil
}

// UpdateUser updates the user specified in received JSON
func (mp *MockUsers) UpdateUser(ctx context.Context, user *data.User) error {
	_, span := otel.Tracer("users").Start(ctx, "updateUserDatabase")
	defer span.End()
	index := findIndexByUserID(user.ID)
	if index == -1 {
		return data.ErrorUserNotFound
	}
	userList[index] = user
	return nil
}

// DeleteUser deletes the user with the given id
func (mp *MockUsers) DeleteUser(ctx context.Context, id string) error {
	_, span := otel.Tracer("users").Start(ctx, "deleteUserDatabase")
	defer span.End()
	index := findIndexByUserID(id)
	if index == -1 {
		return data.ErrorUserNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	userList = append(userList[:index], userList[index+1:]...)

	return nil
}

// Returns the index of a user in the database
// Returns -1 when no user is found
func findIndexByUserID(id string) int {
	for index, user := range userList {
		if user.ID == id {
			return index
		}
	}
	return -1
}

// Returns the index of a user in the database
// Returns -1 when no user is found
func findIndexByUsername(username string) int {
	for index, user := range userList {
		if user.Username == username {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///////////////////////////////////////////////////////////////////////////

var userList = []*data.User{
	{
		ID:          "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Username:    "JeremiS",
		Email:       "jeremi@gmail.com",
		FirstName:   "Jeremi",
		LastName:    "Savard",
		DateOfBirth: "08/02/1996",
		Gender:      "M",
	},
	{
		ID:          "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Username:    "MalcolmSJ",
		Email:       "malcolmb@gmail.com",
		FirstName:   "Malcolm",
		LastName:    "St-John",
		DateOfBirth: "01/01/1994",
		Gender:      "M",
	},
}
