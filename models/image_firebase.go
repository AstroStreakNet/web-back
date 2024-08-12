package models

import (
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
