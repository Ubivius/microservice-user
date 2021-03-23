package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/data"
	"github.com/prometheus/common/log"
)

// Delete a user with specified id from the database
func (userHandler *UsersHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)
	log.Info("Delete user by ID request", "id", id)

	err := userHandler.db.DeleteUser(id)
	switch err {
	case nil:
		responseWriter.WriteHeader(http.StatusNoContent)
		return
	case data.ErrorUserNotFound:
		log.Error(err, "Error deleting user, id does not exist")
		http.Error(responseWriter, "User not found", http.StatusNotFound)
		return
	default:
		log.Error(err, "Error deleting user")
		http.Error(responseWriter, "Error deleting user", http.StatusInternalServerError)
		return
	}
}
