package repository

import (
	"github.com/terrylin13/gin-restful-generator/example/internal/model"
)

// type UserRepository interface {
// 	CreateUser(user *model.User) error
// 	GetUser(user *model.User, id string) error
// }

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUser(user *model.User, id string) error
	GetAllUsers() ([]*model.User, error)
}
