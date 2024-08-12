package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string
	Password    string
	DisplayName string
	FirstName   string
	LastName    string
	Role        string
	Images      []Image // Necessary for HAS MANY relationship in GORM
}
