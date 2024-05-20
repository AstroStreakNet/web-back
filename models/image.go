// models/image.go

package models

import (
	"encoding/json"
	"time"

	"webback/config"
	"webback/middlewares"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type streakType string

const ()

var db *gorm.DB

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

// Initialize the database connection and auto-migrate the Image model
func init() {
	config.Connect()
	db.AutoMigrate(&Image{})
}

func CreateImage(image *Image) *Image {
	db.Create(&image)
	return image
}

func GetAllImages() []Image {
	var images []Image
	db.Find(&images)
	return images
}

func GetImageByID(ID string) (*Image, error) {
	var image Image
	if err := db.First(&image, "image_id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func DeleteImage(ID string) error {
	if err := db.Delete(&Image{}, "image_id = ?", ID).Error; err != nil {
		return err
	}
	return nil
}

func (image *Image) Update(db *gorm.DB) error {
	return db.Save(image).Error
}

func GetAllImagesPublic() []Image {
	var images []Image
	db.Where("allow_public = ?", true).Find(&images)
	return images
}

func GetAllImagesML() []Image {
	var images []Image
	db.Where("allow_ml = ?", true).Find(&images)
	return images
}
