package models

import (
	"gorm.io/datatypes"
	"time"
)

type streakType string

const ()

type Image struct {
	ID                string         `gorm:"primaryKey; column:image_id; type:VARCHAR(45)"`
	ImagePath         string         `gorm:"column:image_path; type:VARCHAR(45)"`
	ObservatoryCode   string         `gorm:"column:observatory_code; type:VARCHAR(45)"`
	TimeOfObservation time.Time      `gorm:"column:time_of_observation; type:DATETIME"`
	ExposureDuration  datatypes.Time `gorm:"column:exposure_duration; type:TIME"`
	// Coordinates
	StreakType streakType `gorm:"column:streak_type; type:ENUM()"`
	UserID     string     `gorm:"foreignKey:user_id; column:fk_user_id; type:VARCHAR(45)"`
}
