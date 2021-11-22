package data

import (
	"fmt"
)

var ErrorUserNotFound = fmt.Errorf("User not found")

// StatusType of a status
type StatusType string

// status type of a player
const (
	Online  StatusType = "Online"  // user is online
	Offline StatusType = "Offline" // user is offline
	InLobby StatusType = "InLobby" // user is in lobby
	InGame  StatusType = "InGame"  // user is in game
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
