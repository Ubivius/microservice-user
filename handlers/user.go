package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/Ubivius/user-microservice/data"
)

type User struct {
	l *log.Logger
}

func NewUser(l *log.Logger) *User {
	return &User{l}
}

func (u *User) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lu := data.GetUsers()
	d, err := json.Marshal(lu)
	if err!=nil {
		http.Error(rw, "Unable to marshal json",http.StatusInternalServerError)
	}

	rw.Write(d)
}
