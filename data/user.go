package data

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	CreatedOn string `json:"-"`
}

// Users is a collection of User
type Users []*User

//Return every users
func GetUsers() Users {
	return userList
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
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
