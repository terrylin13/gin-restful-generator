package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}
