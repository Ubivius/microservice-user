package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/data"
)

// MiddlewareUserValidation is used to validate incoming user JSONS
func (userHandler *UsersHandler) MiddlewareUserValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		user := &data.User{}

		err := json.NewDecoder(request.Body).Decode(user)
		if err != nil {
			http.Error(responseWriter, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		//validate the user
		err = user.Validate()
		if err != nil {
			userHandler.logger.Println("[ERROR] validating user", err)
			http.Error(
				responseWriter,
				fmt.Sprintf("Error validating user: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		//Add the user to the context
		context := context.WithValue(request.Context(), KeyUser{}, user)
		newRequest := request.WithContext(context)

		// Call the nest handler
		next.ServeHTTP(responseWriter, newRequest)
	})
}
