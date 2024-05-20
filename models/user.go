// models/user.go

package models

import (
	"errors"

	"gorm.io/gorm"
)

// Initialize DB from the config package
var DB *gorm.DB

type User struct {
	ID           string `gorm:"primaryKey; column:user_id; type:VARCHAR(45)"`
	AstronomerID uint   `gorm:"column:astronomer_id; type: INT"`
	Password     string `gorm:"column:#password; type:VARCHAR(45)"`
	FirstName    string `gorm:"column:first_name; type:VARCHAR(45)"`
	LastName     string `gorm:"column:last_name; type:VARCHAR(45)"`
	DOB          string `gorm:"column:dob; type:DATE"`
	Gender       string `gorm:"column:gender; type:ENUM('MALE', 'FEMALE', 'UNDEFINED')"`
	Permissions  uint   `gorm:"column:permissions; type:TINYINT"`
	ImagePublish uint   `gorm:"column:image_publish; type:TINYINT"`
}

// CreateUser creates a new user
func CreateUser(user *User) (*User, error) {
	result := DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetAllUsers fetches all users
func GetAllUsers() ([]User, error) {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUserByID fetches a user by ID
func GetUserByID(id string) (*User, error) {
	var user User
	result := DB.First(&user, "user_id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByAstronomerID fetches a user by Astronomer ID
func GetUserByAstronomerID(astronomerID int) (*User, error) {
	var user User
	result := DB.First(&user, "astronomer_id = ?", astronomerID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// DeleteUser deletes a user by ID
func DeleteUser(id string) error {
	result := DB.Delete(&User{}, "user_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// UpdateUser updates a user's details
func UpdateUser(user *User) (*User, error) {
	result := DB.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
