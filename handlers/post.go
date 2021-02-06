package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/data"
)

// AddProduct creates a new user from the received JSON
func (userHandler *UsersHandler) AddUser(responseWriter http.ResponseWriter, request *http.Request) {
	userHandler.logger.Println("Handle POST User")

	user := request.Context().Value(KeyUser{}).(*data.User)
	data.AddUser(user)
}
