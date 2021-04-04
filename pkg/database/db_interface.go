package database

import (
	"github.com/Ubivius/microservice-user/pkg/data"
)

// The interface that any kind of database must implement
type UserDB interface {
	GetUsers() data.Users
	GetUserByID(id string) (*data.User, error)
	UpdateUser(user *data.User) error
	AddUser(user *data.User) error
	DeleteUser(id string) error
	Connect() error
	CloseDB()
}
