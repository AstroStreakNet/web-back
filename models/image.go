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
	AllowPublic       bool           `gorm:"column:allow_public; type:TINYINT"` // Addition, not in Emily schema
	AllowML           bool           `gorm:"column:allow_ml; type:TINYINT"`     // Addition, not in Emily schema
	ObservatoryCode   string         `gorm:"column:observatory_code; type:VARCHAR(45)"`
	TimeOfObservation time.Time      `gorm:"column:time_of_observation; type:DATETIME"`
	ExposureDuration  datatypes.Time `gorm:"column:exposure_duration; type:TIME"`
	StreakType        streakType     `gorm:"column:streak_type; type:ENUM()"`
	UserID            string         `gorm:"foreignKey:user_id; column:fk_user_id; type:VARCHAR(45)"`
	// Coordinates, need to change datatype, POINT won't work
	// TODO add json tags
}

// MarshalJSON custom JSON marshaling for public image response
func (i Image) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON custom JSON unmarshalling for image post request
func (i Image) UnmarshalJSON(data []byte) error {
	return nil
}
