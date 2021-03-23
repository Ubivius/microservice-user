package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/data"
)

// UpdateUsers updates the user with the ID specified in the received JSON user
func (userHandler *UsersHandler) UpdateUsers(responseWriter http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(KeyUser{}).(*data.User)
	log.Info("UpdateUsers request", "id", user.ID)

	err := userHandler.db.UpdateUser(user)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorUserNotFound:
		log.Error(err, "User not found")
		http.Error(responseWriter, "User not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error updating user")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
