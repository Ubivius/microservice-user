package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/data"
)

// AddUser creates a new user from the received JSON
func (userHandler *UsersHandler) AddUser(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("AddUser request")

	user := request.Context().Value(KeyUser{}).(*data.User)
	err := userHandler.db.AddUser(user)

	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	default:
		log.Error(err, "Error adding user")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
