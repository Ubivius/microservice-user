package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ubivius/microservice-user/handlers"
	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
)

func main() {
	//Logger
	logger := log.New(os.Stdout, "microservice-user", log.LstdFlags)

	// Configuration elastic search
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logger.Println("Error creating the es client.")
	}
	var b bytes.Buffer
	b.WriteString(`{"Users" : "Jeremi"}`)

	res, _ := es.Index("method1", &b)
	logger.Println(res)

	// Handlers
	userHandler := handlers.NewUsersHandler(logger)

	// Mux route handling with gorilla/mux
	router := mux.NewRouter()

	//Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/users", userHandler.GetUsers)
	getRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID)

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
	deleteRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.Delete)

	// Start server
	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     router,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port ", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			logger.Println("Error starting server : ", err)
			logger.Fatal(err)
		}
	}()

	// Handle shutdown signals from operating system
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	receivedSignal := <-signalChannel

	logger.Println("Received terminate, beginning graceful shutdown", receivedSignal)

	// Server shutdown
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_ = server.Shutdown(timeoutContext)
}
