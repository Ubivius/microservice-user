package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-user/data"
)

// MiddlewareUserValidation is used to validate incoming user JSONS
func (userHandler *UsersHandler) MiddlewareUserValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		user := &data.User{}

		err := data.FromJSON(user, request.Body)
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
