// models/image.go

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
	ID string `gorm:"primaryKey; column:image_id; type:VARCHAR(45)"`
	// MetaData
	ImagePath   string `gorm:"column:image_path; type:VARCHAR(45)"`
	AllowPublic bool   `gorm:"column:allow_public; type:TINYINT"`
	AllowML     bool   `gorm:"column:allow_ml; type:TINYINT"`
	// Coordinates
	RightAscension   float32        `gorm:"column:right_ascension; type:FLOAT"`
	Declination      float32        `gorm:"column:declination; type:FLOAT"`
	JulianDate       time.Time      `gorm:"column:time_of_observation; type:DATETIME"`
	ExposureDuration datatypes.Time `gorm:"column:exposure_duration; type:TIME"`
	// Observation data
	ObservatoryCode string     `gorm:"column:observatory_code; type:VARCHAR(45)"`
	StreakType      streakType `gorm:"column:streak_type; type:ENUM()"`
	UserID          string     `gorm:"foreignKey:user_id; column:fk_user_id; type:VARCHAR(45)"`
	// Astrometry references
	AstroSubID     int  `gorm:"column:astro_sub_id; type:VARCHAR(45)"`
	AstroProcessed bool `gorm:"column:astro_processed; type:TINYINT"`
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
		UploadDateTime: i.JulianDate.Unix(),
		StaticPath:     middlewares.GetStaticURL(i.ID),
		Tags:           []string{"astronomy"}, // Need to get tags into the db
	}
	return json.Marshal(publicImage)
}
