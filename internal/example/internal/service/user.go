package service

import (
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (service *UserService) CreateUser(user *model.User) error {
	return service.repo.CreateUser(user)
}

func (service *UserService) GetUser(user *model.User, id string) error {
	return service.repo.GetUser(user, id)
}

func (service *UserService) GetAllUsers() ([]*model.User, error) {
	return service.repo.GetAllUsers()
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
