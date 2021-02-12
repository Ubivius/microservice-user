package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/data"
)

// GetUsers returns the full list of users
func (userHandler *UsersHandler) GetUsers(responseWriter http.ResponseWriter, request *http.Request) {
	userHandler.logger.Println("Handle GET Users")

	// fetch the users from the datastore
	userList := data.GetUsers()

	// serialize the list to JSON
	err := data.ToJSON(userList, responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetUserByID returns a single user from the database
func (userHandler *UsersHandler) GetUserByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)

	userHandler.logger.Println("[DEBUG] getting id", id)

	user, err := data.GetUserByID(id)
	switch err {
	case nil:
	case data.ErrorUserNotFound:
		userHandler.logger.Println("[ERROR] fetching user", err)
		http.Error(responseWriter, "User not found", http.StatusBadRequest)
		return
	default:
		userHandler.logger.Println("[ERROR] fetching user", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(user, responseWriter)
	if err != nil {
		userHandler.logger.Println("[ERROR] serializing user", err)
	}
}
