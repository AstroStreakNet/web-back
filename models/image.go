package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserID           *uint // ForeignKey, necessary for HAS MANY relationship in gorm
	Path             string
	URL              string
	AstrometryID     *string
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
	userID           *uint
	path             string
	url              string
	astrometryID     *string
	observatoryCode  string
	rightAscension   string
	declination      string
	julianDate       string
	exposureDuration string
	streakType       string
	tags             string
	allowPublic      bool
	allowML          bool
}

func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{}
}

func (builder *ImageBuilder) WithUserID(ID *uint) *ImageBuilder {
	builder.userID = ID
	return builder
}

func (builder *ImageBuilder) WithPath(path string) *ImageBuilder {
	builder.path = path
	return builder
}

func (builder *ImageBuilder) WithURL(url string) *ImageBuilder {
	builder.url = url
	return builder
}

func (builder *ImageBuilder) WithAstrometryID(astrometryID string) *ImageBuilder {
	builder.astrometryID = &astrometryID
	return builder
}

func (builder *ImageBuilder) WithObservatoryCode(observatoryCode string) *ImageBuilder {
	builder.observatoryCode = observatoryCode
	return builder
}

func (builder *ImageBuilder) WithRightAscension(rightAscension string) *ImageBuilder {
	builder.rightAscension = rightAscension
	return builder
}

func (builder *ImageBuilder) WithDeclination(declination string) *ImageBuilder {
	builder.declination = declination
	return builder
}

func (builder *ImageBuilder) WithJulianDate(julianDate string) *ImageBuilder {
	builder.julianDate = julianDate
	return builder
}

func (builder *ImageBuilder) WithExposureDuration(exposureDuration string) *ImageBuilder {
	builder.exposureDuration = exposureDuration
	return builder
}

func (builder *ImageBuilder) WithStreakType(streakType string) *ImageBuilder {
	builder.streakType = streakType
	return builder
}

func (builder *ImageBuilder) WithTags(tags string) *ImageBuilder {
	builder.tags = tags
	return builder
}

func (builder *ImageBuilder) WithAllowPublic(allowPublic bool) *ImageBuilder {
	builder.allowPublic = allowPublic
	return builder
}

func (builder *ImageBuilder) WithAllowML(allowML bool) *ImageBuilder {
	builder.allowML = allowML
	return builder
}

func (builder *ImageBuilder) Build() *Image {
	return &Image{
		UserID:           builder.userID,
		Path:             builder.path,
		URL:              builder.url,
		AstrometryID:     builder.astrometryID,
		ObservatoryCode:  builder.observatoryCode,
		RightAscension:   builder.rightAscension,
		Declination:      builder.declination,
		JulianDate:       builder.julianDate,
		ExposureDuration: builder.exposureDuration,
		StreakType:       builder.streakType,
		Tags:             builder.tags,
		AllowPublic:      builder.allowPublic,
		AllowML:          builder.allowML,
	}
}
