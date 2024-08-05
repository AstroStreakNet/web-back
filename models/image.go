// models/image.go

package models

import (
	"gorm.io/gorm"
	"time"
)

type Image struct {
	gorm.Model
	// ForeignKey, necessary for HAS MANY relationship in gorm
	UserID *uint
	// Path
	Path         string
	URL          *string
	AstrometryID *string
	// Image metadata
	ObservatoryCode  *string
	RightAscension   *string
	Declination      *string
	JulianDate       *time.Time
	ExposureDuration *time.Time
	StreakType       *string
	Tags             *string
	AllowPublic      bool
	AllowML          bool
}
