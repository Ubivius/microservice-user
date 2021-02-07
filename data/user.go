package data

import (
	"fmt"
)

var ErrorUserNotFound = fmt.Errorf("User not found")

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	FisrtName    string `json:"firstname"`
	Name         string `json:"name"`
	DateOfBirth  string `json:"dateofbirth" validate:"required,dateofbirth"`
	Gender       string `json:"gender"`
	Address      string `json:"address"`
	Bio          string `json:"bio"`
	Achievements string `json:"achievements"`
}

// Users is a collection of User
type Users []*User

//Return every users
func GetUsers() Users {
	return userList
}

// GetUserByID returns a single product with the given id
func GetUserByID(id int) (*User, error) {
	index := findIndexByUserID(id)
	if id == -1 {
		return nil, ErrorUserNotFound
	}
	return userList[index], nil
}

// AddUser creates a new user
func AddUser(user *User) {
	user.ID = getNextID()
	userList = append(userList, user)
}

// UpdateUser updates the user specified in received JSON
func UpdateUser(user *User) error {
	index := findIndexByUserID(user.ID)
	if index == -1 {
		return ErrorUserNotFound
	}
	userList[index] = user
	return nil
}

// DeleteUser deletes the user with the given id
func DeleteUser(id int) error {
	index := findIndexByUserID(id)
	if index == -1 {
		return ErrorUserNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	userList = append(userList[:index], userList[index+1])

	return nil
}

// Returns the index of a user in the database
// Returns -1 when no user is found
func findIndexByUserID(id int) int {
	for index, user := range userList {
		if user.ID == id {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

func getNextID() int {
	userList := userList[len(userList)-1]
	return userList.ID + 1
}

var userList = []*User{
	{
		ID:       1,
		Username: "sickboy",
	},
	{
		ID:       2,
		Username: "Mark Renton",
	},
}
