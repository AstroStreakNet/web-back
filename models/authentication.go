// models/authentication.go

// mdoel copied over from database API

package models

import (
	"time"

	"webback/config"
)

type Authentication struct {
	ID           uint      `gorm:"primaryKey; column:auth_id; type:INT"`
	Timestamp    time.Time `gorm:"column:timestamp; type:DATETIME"`
	Success      bool      `gorm:"column:success; type:TINYINT"` // Booleans get converted to TINYINT through Gorm
	UserID       string    `gorm:"foreignKey:user_id; column:fk_user_id; type:VARCHAR(45)"`
	LoginTime    time.Time `gorm:"column:login_time; type:DATETIME"`
	LogoutTime   time.Time `gorm:"column:logout_time; type:DATETIME"`
	SessionToken string    `gorm:"column:session_token; type:VARCHAR(255)"`
	// TODO: Ask Emily about Schema
}

func init() {
	config.Connect()
	db.AutoMigrate(&Authentication{})
}

func CreateAuthentication(auth *Authentication) *Authentication {
	db.Create(auth)
	return auth
}

func GetAllAuthentications() []Authentication {
	var authentications []Authentication
	db.Find(&authentications)
	return authentications
}

func GetAuthenticationByID(ID uint) (*Authentication, error) {
	var auth Authentication
	if err := db.First(&auth, ID).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func DeleteAuthentication(ID uint) error {
	if err := db.Delete(&Authentication{}, ID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAuthentication(auth *Authentication) (*Authentication, error) {
	if err := db.Save(auth).Error; err != nil {
		return nil, err
	}
	return auth, nil
}
