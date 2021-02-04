package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Ubivius/user-microservice/data"
	"github.com/gorilla/mux"
)

type Users struct {
	l *log.Logger
}

type KeyUser struct{}

// NewUsers creates a products handler with the given logger
func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
/*func (u *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of Users
	if r.Method == http.MethodGet {
		u.GetUsers(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		u.AddUser(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		//expect the Id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			u.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			u.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			u.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		u.l.Println("got id", id)

		u.UpdateUsers(rw, r)
	}

	// catch all
	// if no method is satisfied return an error
	//rw.WriteHeader(http.StatusMethodNotAllowed)
}*/

// getProducts returns the products from the data store
func (u *Users) GetUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle GET Users")

	// fetch the products from the datastore
	lu := data.GetUsers()

	// serialize the list to JSON
	err := lu.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (u *Users) AddUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handle POST User")

	user := r.Context().Value(KeyUser{}).(data.User)
	data.AddUser(&user)
}

func (u Users) UpdateUsers(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(KeyUser{}).(data.User)
	u.l.Println("Handle PUT User", id)

	err = data.UpdateUser(id, &user)
	if err == data.ErrUserNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "User not found", http.StatusInternalServerError)
		return
	}
}

func (u Users) MiddlewareUserValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := data.User{}

		err := user.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		//validate the product
		err = user.Validate()
		if err != nil {
			u.l.Println("[ERROR] validating user", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating user: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
