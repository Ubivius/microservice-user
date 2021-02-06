package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/data"
)

// Delete a product with specified id from the database
func (userHandler *UsersHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getUserID(request)
	userHandler.logger.Println("Handle DELETE user", id)

	err := data.DeleteUser(id)
	if err == data.ErrorProductNotFound {
		userHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		userHandler.logger.Println("[ERROR] deleting user", err)
		http.Error(responseWriter, "Erro deleting user", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
