package models

import (
	"encoding/json"
	"gorm.io/datatypes"
	"time"
	"webback/middlewares"
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
	publicImage := struct {
		Name           string   `json:"name"`
		Uploader       string   `json:"uploader"`
		UploadDateTime int64    `json:"uploadDate"`
		StaticPath     string   `json:"url"`
		Tags           []string `json:"tags"`
	}{
		Name:           i.ID,    // Probably needs to be replaced
		Uploader:       "Steve", // Sort of unnecessary, will have to discuss with Adrian
		UploadDateTime: i.TimeOfObservation.Unix(),
		StaticPath:     middlewares.GetStaticURL(i.ID),
		Tags:           []string{"astronomy"}, // Need to get tags into the db
	}
	return json.Marshal(publicImage)
}

// UnmarshalJSON custom JSON unmarshalling for image post request
func (i Image) UnmarshalJSON(data []byte) error {
	return nil
}
