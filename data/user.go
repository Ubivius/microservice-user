package data

import (
	"encoding/json"
	"fmt"
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

func AddUser(u *User) {
	u.Id = getNextID()
	userList = append(userList, u)
}

func UpdateUser(id int, u *User) error {
	_, pos, err := findUser(id)
	if err != nil {
		return err
	}

	u.Id = id
	userList[pos] = u

	return nil
}

var ErrUserNotFound = fmt.Errorf("User not found")

func findUser(id int) (*User, int, error) {
	for i, u := range userList {
		if u.Id == id {
			return u, i, nil
		}
	}
	return nil, -1, ErrUserNotFound
}

func getNextID() int {
	lu := userList[len(userList)-1]
	return lu.Id + 1
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

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

var userList = []*User{
	{
		Id:        12345,
		Username:  "sickboy",
		CreatedOn: time.Now().UTC().String(),
	},
	{
		Id:        54321,
		Username:  "Mark Renton",
		CreatedOn: time.Now().UTC().String(),
	},
}
