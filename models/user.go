package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name     string
	email    string
	password string // hashed
}
