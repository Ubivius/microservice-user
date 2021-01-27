package handlers

import (
	"log"
	"net/http"

	"github.com/Ubivius/user-microservice/data"
)

type Users struct {
	l *log.Logger
}

// NewUsers creates a products handler with the given logger
func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of Users
	if r.Method == http.MethodGet {
		u.getUsers(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (u *Users) getUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Users")

	// fetch the products from the datastore
	lu := data.GetUsers()

	// serialize the list to JSON
	err := lu.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
