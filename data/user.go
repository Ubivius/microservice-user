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
	DateOfBirth  string `json:"dateofbirth"`
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

// GetProductByID returns a single product with the given id
func GetUserByID(id int) (*User, error) {
	index := findIndexByUserID(id)
	if id == -1 {
		return nil, ErrorUserNotFound
	}
	return userList[index], nil
}

func AddUser(user *User) {
	user.ID = getNextID()
	userList = append(userList, user)
}

func UpdateUser(user *User) error {
	index := findIndexByUserID(user.ID)
	if index == -1 {
		return ErrorUserNotFound
	}
	userList[index] = user
	return nil
}

func findIndexByUserID(id int) int {
	for index, user := range userList {
		if user.ID == id {
			return index
		}
	}
	return -1
}

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
