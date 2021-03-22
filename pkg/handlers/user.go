package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UsersHandler contains the items common to all user handler functions
type UsersHandler struct {
	logger *log.Logger
}

// KeyUser is a key used for the User object inside context
type KeyUser struct{}

// NewUsersHandler creates a user handler with the given logger
func NewUsersHandler(logger *log.Logger) *UsersHandler {
	return &UsersHandler{logger}
}

// getUserID extracts the user ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getUserID(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
