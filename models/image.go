package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserID           *uint // ForeignKey, necessary for HAS MANY relationship in gorm
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

// Builder

type ImageBuilder struct {
	UserID           *uint
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

func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{}
}

func (builder *ImageBuilder) WithUserID(ID *uint) *ImageBuilder {
	builder.UserID = ID
	return builder
}

func (builder *ImageBuilder) WithPath(path string) *ImageBuilder {
	builder.Path = path
	return builder
}

func (builder *ImageBuilder) WithURL(url string) *ImageBuilder {
	builder.URL = url
	return builder
}

func (builder *ImageBuilder) WithAstrometryID(astrometryID string) *ImageBuilder {
	builder.AstrometryID = astrometryID
	return builder
}

func (builder *ImageBuilder) WithObservatoryCode(observatoryCode string) *ImageBuilder {
	builder.ObservatoryCode = observatoryCode
	return builder
}

func (builder *ImageBuilder) WithRightAscension(rightAscension string) *ImageBuilder {
	builder.RightAscension = rightAscension
	return builder
}

func (builder *ImageBuilder) WithDeclination(declination string) *ImageBuilder {
	builder.Declination = declination
	return builder
}

func (builder *ImageBuilder) WithJulianDate(julianDate string) *ImageBuilder {
	builder.JulianDate = julianDate
	return builder
}

func (builder *ImageBuilder) WithExposureDuration(exposureDuration string) *ImageBuilder {
	builder.ExposureDuration = exposureDuration
	return builder
}

func (builder *ImageBuilder) WithStreakType(streakType string) *ImageBuilder {
	builder.StreakType = streakType
	return builder
}

func (builder *ImageBuilder) WithTags(tags string) *ImageBuilder {
	builder.Tags = tags
	return builder
}

func (builder *ImageBuilder) WithAllowPublic(allowPublic bool) *ImageBuilder {
	builder.AllowPublic = allowPublic
	return builder
}

func (builder *ImageBuilder) WithAllowML(allowML bool) *ImageBuilder {
	builder.AllowML = allowML
	return builder
}

func (builder *ImageBuilder) Build() *Image {
	return &Image{
		UserID:           builder.UserID,
		Path:             builder.Path,
		URL:              builder.URL,
		AstrometryID:     builder.AstrometryID,
		ObservatoryCode:  builder.ObservatoryCode,
		RightAscension:   builder.RightAscension,
		Declination:      builder.Declination,
		JulianDate:       builder.JulianDate,
		ExposureDuration: builder.ExposureDuration,
		StreakType:       builder.StreakType,
		Tags:             builder.Tags,
		AllowPublic:      builder.AllowPublic,
		AllowML:          builder.AllowML,
	}
}
