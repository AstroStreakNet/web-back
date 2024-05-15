// models/authentication.go

package models

import (
    "time"
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

