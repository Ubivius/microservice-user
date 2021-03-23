package data

import (
	"fmt"
)

var ErrorUserNotFound = fmt.Errorf("User not found")

type User struct {
	ID           string `json:"id" bson:"_id"`
	Username     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"name"`
	DateOfBirth  string `json:"dateofbirth" validate:"required,dateofbirth"`
	Gender       string `json:"gender"`
	Address      string `json:"address"`
	Bio          string `json:"bio"`
	Achievements string `json:"achievements"`
}

// Users is a collection of User
type Users []*User
