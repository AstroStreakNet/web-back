package models

import "time"

type UserGorm struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	Email       string
	Password    string
	DisplayName string
	FirstName   string
	LastName    string
	Role        string
	Images      []ImageGorm // Necessary for HAS MANY relationship in GORM
}

func UserFromGorm(gorm UserGorm) User {
	return User{
		ID:          gorm.ID,
		CreatedAt:   gorm.CreatedAt,
		Email:       gorm.Email,
		DisplayName: gorm.DisplayName,
		FirstName:   gorm.FirstName,
		LastName:    gorm.LastName,
		Role:        gorm.Role,
	}
}

func GormFromUser(user User) UserGorm {
	return UserGorm{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Role:        user.Role,
	}
}
