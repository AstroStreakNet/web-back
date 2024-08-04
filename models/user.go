// models/user.go

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string
	Password    string
	DisplayName string
	FirstName   *string
	LastName    *string
	Role        string
	// This is necessary to create the HAS MANY relationship in gorm.
	Images []Image
}
