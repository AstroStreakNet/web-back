package models

import (
	"strconv"
	"strings"
	"time"
)

type ImageFirebase struct {
	ID               string    `firestore:"Document ID"`
	CreatedAt        time.Time `firestore:"created_at"`
	User             string    `firestore:"user"`
	Path             string    `firestore:"path"`
	URL              string    `firestore:"url"`
	AstrometryID     string    `firestore:"astrometry_id"`
	ObservatoryCode  string    `firestore:"observatory_code"`
	RightAscension   string    `firestore:"right_ascension"`
	Declination      string    `firestore:"declination"`
	JulianDate       string    `firestore:"julian_date"`
	ExposureDuration string    `firestore:"exposure_duration"`
	StreakType       string    `firestore:"streak_type"`
	Tags             []string  `firestore:"tags"`
	AllowPublic      bool      `firestore:"allow_public"`
	AllowML          bool      `firestore:"allow_ml"`
}

func imageFromFirebase(firebase ImageFirebase) *Image {

	// User reference to id
	userID64, err := strconv.ParseUint(firebase.User, 10, 32)
	if err != nil {
		//
	}
	userID := uint(userID64)

	// Tags string array to single string
	tags := strings.Join(firebase.Tags, " ")

	return &Image{
		CreatedAt:        firebase.CreatedAt,
		UserID:           userID,
		Path:             firebase.Path,
		URL:              firebase.URL,
		AstrometryID:     firebase.AstrometryID,
		ObservatoryCode:  firebase.ObservatoryCode,
		RightAscension:   firebase.RightAscension,
		Declination:      firebase.Declination,
		JulianDate:       firebase.JulianDate,
		ExposureDuration: firebase.ExposureDuration,
		StreakType:       firebase.StreakType,
		Tags:             tags,
		AllowPublic:      firebase.AllowPublic,
		AllowML:          firebase.AllowML,
	}
}
