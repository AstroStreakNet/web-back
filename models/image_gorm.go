package models

import (
	"time"
)

type ImageGorm struct {
	ID               uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	UserGormID       *uint // ForeignKey, necessary for HAS MANY relationship in gorm
	Path             string
	URL              string
	AstrometryID     string
	ObservatoryCode  string
	RightAscension   string
	Declination      string
	JulianDate       string
	ExposureDuration string
	StreakType       string
	Tags             string
	AllowPublic      bool
	AllowML          bool
}

func ImageFromGorm(gorm ImageGorm) Image {

	var userID *uint
	if gorm.UserGormID == nil {
		userID = new(uint)
		*userID = 0
	} else {
		userID = gorm.UserGormID
	}

	return Image{
		ID:               gorm.ID,
		CreatedAt:        gorm.CreatedAt,
		UserID:           *userID,
		Path:             gorm.Path,
		URL:              gorm.URL,
		AstrometryID:     gorm.AstrometryID,
		ObservatoryCode:  gorm.ObservatoryCode,
		RightAscension:   gorm.RightAscension,
		Declination:      gorm.Declination,
		JulianDate:       gorm.JulianDate,
		ExposureDuration: gorm.ExposureDuration,
		StreakType:       gorm.StreakType,
		Tags:             gorm.Tags,
		AllowPublic:      gorm.AllowPublic,
		AllowML:          gorm.AllowML,
	}
}

func GormFromImage(image Image) ImageGorm {
	return ImageGorm{
		ID:               image.ID,
		CreatedAt:        image.CreatedAt,
		UserGormID:       &image.UserID,
		Path:             image.Path,
		URL:              image.URL,
		AstrometryID:     image.AstrometryID,
		ObservatoryCode:  image.ObservatoryCode,
		RightAscension:   image.RightAscension,
		Declination:      image.Declination,
		JulianDate:       image.JulianDate,
		ExposureDuration: image.ExposureDuration,
		StreakType:       image.StreakType,
		Tags:             image.Tags,
		AllowPublic:      image.AllowPublic,
		AllowML:          image.AllowML,
	}
}
