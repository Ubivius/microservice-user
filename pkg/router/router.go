package router

import (
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/handlers"
	"github.com/gorilla/mux"
)

// Mux route handling with gorilla/mux
func New(userHandler *handlers.UsersHandler) *mux.Router {
	// Mux route handling with gorilla/mux
	log.Info("Starting router")
	router := mux.NewRouter()

	//Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/users", userHandler.GetUsers)
	getRouter.HandleFunc("/users/{id:[0-9a-z-]+}", userHandler.GetUserByID)

	//Put Router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/users", userHandler.UpdateUsers)
	putRouter.Use(userHandler.MiddlewareUserValidation)

	//Post Router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/users", userHandler.AddUser)
	postRouter.Use(userHandler.MiddlewareUserValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/users/{id:[0-9a-z-]+}", userHandler.Delete)

	return router
}
