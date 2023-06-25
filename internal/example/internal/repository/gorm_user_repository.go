package repository

import (
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/config"
	"github.com/terrylin13/gin-restful-generator/internal/example/internal/model"
)

type GormUserRepository struct {
}

func (repo *GormUserRepository) CreateUser(user *model.User) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	return db.Create(user).Error
}

func (repo *GormUserRepository) GetUser(user *model.User, id string) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	return db.First(user, id).Error
}

func (repo *GormUserRepository) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
