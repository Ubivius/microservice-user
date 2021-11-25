package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/data"
	"go.opentelemetry.io/otel"
)

// AddUser creates a new user from the received JSON
func (userHandler *UsersHandler) AddUser(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("user").Start(request.Context(), "addUser")
	defer span.End()
	log.Info("AddUser request")

	user := request.Context().Value(KeyUser{}).(*data.User)
	err := userHandler.db.AddUser(request.Context(), user)

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
