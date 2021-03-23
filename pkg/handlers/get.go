package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/data"
	"github.com/prometheus/common/log"
)

// GetUsers returns the full list of users
func (userHandler *UsersHandler) GetUsers(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("GetUsers request")
	userList := userHandler.db.GetUsers()
	err := json.NewEncoder(responseWriter).Encode(userList)
	if err != nil {
		log.Error(err, "Error serializing user")
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetUserByID returns a single user from the database
func (userHandler *UsersHandler) GetUserByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)

	log.Info("GetUserByID request", "id", id)

	user, err := userHandler.db.GetUserByID(id)

	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(user)
		if err != nil {
			log.Error(err, "Error serializing user")
		}
		return
	case data.ErrorUserNotFound:
		log.Error(err, "User not found")
		http.Error(responseWriter, "user not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error getting user")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
