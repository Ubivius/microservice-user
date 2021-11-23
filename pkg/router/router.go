package router

import (
	"net/http"

	"github.com/Ubivius/microservice-user/pkg/handlers"
	"github.com/Ubivius/pkg-telemetry/metrics"
	tokenValidation "github.com/Ubivius/shared-authentication/pkg/auth"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

// Mux route handling with gorilla/mux
func New(userHandler *handlers.UsersHandler) *mux.Router {
	// Mux route handling with gorilla/mux
	log.Info("Starting router")
	router := mux.NewRouter()
	router.Use(otelmux.Middleware("users"))
	router.Use(metrics.RequestCountMiddleware)

	//Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.Use(tokenValidation.Middleware)
	getRouter.HandleFunc("/users", userHandler.GetUsers)
	getRouter.HandleFunc("/users/{id:[0-9a-z-]+}", userHandler.GetUserByID)

	//Health Check
	healthRouter := router.Methods(http.MethodGet).Subrouter()
	healthRouter.HandleFunc("/health/live", userHandler.LivenessCheck)
	healthRouter.HandleFunc("/health/ready", userHandler.ReadinessCheck)

	//Put Router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.Use(tokenValidation.Middleware)
	putRouter.HandleFunc("/users", userHandler.UpdateUsers)
	putRouter.Use(userHandler.MiddlewareUserValidation)

	//Post Router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(tokenValidation.Middleware)
	postRouter.HandleFunc("/users", userHandler.AddUser)
	postRouter.Use(userHandler.MiddlewareUserValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.Use(tokenValidation.Middleware)
	deleteRouter.HandleFunc("/users/{id:[0-9a-z-]+}", userHandler.Delete)

	return router
}
