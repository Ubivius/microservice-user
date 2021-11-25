package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/data"
	"go.opentelemetry.io/otel"
)

// GetUsers returns the full list of users
func (userHandler *UsersHandler) GetUsers(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("user").Start(request.Context(), "getUsers")
	defer span.End()
	log.Info("GetUsers request")
	userList := userHandler.db.GetUsers(request.Context())
	err := json.NewEncoder(responseWriter).Encode(userList)
	if err != nil {
		log.Error(err, "Error serializing user")
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetUserByID returns a single user from the database
func (userHandler *UsersHandler) GetUserByID(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("user").Start(request.Context(), "getUsersById")
	defer span.End()
	id := getUserID(request)

	log.Info("GetUserByID request", "id", id)

	user, err := userHandler.db.GetUserByID(request.Context(), id)

	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(user)
		if err != nil {
			log.Error(err, "Error serializing user")
		}
		return
	case data.ErrorUserNotFound:
		log.Error(err, "User not found")
		http.Error(responseWriter, "User not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error getting user")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetUserByUsername returns a single user from the database
func (userHandler *UsersHandler) GetUserByUsername(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("user").Start(request.Context(), "getUsersByUsername")
	defer span.End()
	username := getUsername(request)

	log.Info("GetUserByUsername request", "username", username)

	user, err := userHandler.db.GetUserByUsername(request.Context(), username)

	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(user)
		if err != nil {
			log.Error(err, "Error serializing user")
		}
		return
	case data.ErrorUserNotFound:
		log.Error(err, "User not found")
		http.Error(responseWriter, "User not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error getting user")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
