package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/database"
	"github.com/gorilla/mux"
)

// UsersHandler contains the items common to all user handler functions
type UsersHandler struct {
	db database.UserDB
}

// KeyUser is a key used for the User object inside context
type KeyUser struct{}

// NewUsersHandler creates a user handler with the given logger
func NewUsersHandler(db database.UserDB) *UsersHandler {
	return &UsersHandler{db}
}

// getUserID extracts the user ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getUserID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]
	return id
}
