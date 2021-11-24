package data

import (
	"fmt"
)

var ErrorUserNotFound = fmt.Errorf("User not found")

// StatusType of a status
type StatusType string

const (
	Online  StatusType = "Online"
	Offline StatusType = "Offline"
	InLobby StatusType = "InLobby"
	InGame  StatusType = "InGame"
)

type User struct {
	ID           string     `json:"id" bson:"_id"`
	Username     string     `json:"username" validate:"required"`
	Email        string     `json:"email" validate:"required,email"`
	FirstName    string     `json:"firstname"`
	LastName     string     `json:"lastname"`
	DateOfBirth  string     `json:"dateofbirth"`
	Status       StatusType `json:"status" validate:"isStatusType"`
	Gender       string     `json:"gender"`
	Address      string     `json:"address"`
	Bio          string     `json:"bio"`
	Achievements string     `json:"achievements"`
}

// Users is a collection of User
type Users []*User
