package database

import (
	"context"

	"github.com/Ubivius/microservice-user/pkg/data"
)

// The interface that any kind of database must implement
type UserDB interface {
	GetUsers(ctx context.Context) data.Users
	GetUserByID(ctx context.Context, id string) (*data.User, error)
	UpdateUser(ctx context.Context, user *data.User) error
	AddUser(ctx context.Context, user *data.User) error
	DeleteUser(ctx context.Context, id string) error
	Connect() error
	PingDB() error
	CloseDB()
}
