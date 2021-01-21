package data

import (
	"time"
)

type User struct {
	Id        int
	Username  string
	CreatedOn string
}

//Return every users
func GetUsers() []*User {
	return userList
}

var userList = []*User{
	&User{
		Id:        12345,
		Username:  "sickboy",
		CreatedOn: time.Now().UTC().String(),
	},
	&User{
		Id:        54321,
		Username:  "Mark Renton",
		CreatedOn: time.Now().UTC().String(),
	},
}