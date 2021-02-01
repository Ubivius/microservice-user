package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/Ubivius/user-microservice/handlers"
	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
)

func main() {
	//Logger
	l := log.New(os.Stdout, "microservice-user", log.LstdFlags)

	// Configuration elastic search
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		l.Println("Error creating the es client.")
	}
	var b bytes.Buffer
	b.WriteString(`{"Users" : "Jeremi"}`)

	res, _ := es.Index("method1", &b)
	l.Println(res)

	// Handlers
	userHandler := handlers.NewUsers(l)
	l.Println("handler done")

	// Routing
	gorillaMux := mux.NewRouter()

	getRouter := gorillaMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", userHandler.GetUsers)

	putRouter := gorillaMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", userHandler.UpdateUsers)
	l.Println("routing done")

	// Start server
	errServe := http.ListenAndServe(":9090", userHandler)
	l.Println("Starting server on port 9090")
	if errServe != nil {
		l.Printf("Error starting server: %s\n", errServe)
		os.Exit(1)
	}
}
