package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/data"
)

// UpdateUsers updates the user with the ID specified in the received JSON product
func (userHandler *UsersHandler) UpdateUsers(responseWriter http.ResponseWriter, request *http.Request) {
	user := request.Context().Value(KeyUser{}).(*data.User)
	userHandler.logger.Println("Handle PUT User", user.ID)

	err := data.UpdateUser(user)
	if err == data.ErrorUserNotFound {
		userHandler.logger.Println("[ERROR} user not found", err)
		http.Error(responseWriter, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(responseWriter, "Update user failed", http.StatusInternalServerError)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}
