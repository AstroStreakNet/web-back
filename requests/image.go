package requests

import "mime/multipart"

type ImagePostJSON struct {
	FileType         string `json:"file_type"`
	StreakType       string `json:"streak_type"`
	ObservatoryCode  string `json:"observatory_code"`
	RightAscension   string `json:"right_ascension"`
	Declination      string `json:"declination"`
	JulianDate       string `json:"julian_date"`
	ExposureDuration string `json:"exposure_duration"`
	AllowPublic      bool   `json:"allow_public"`
	AllowML          bool   `json:"allow_ml"`
}

type ImagePost struct {
	FileData multipart.FileHeader `form:"file"`
	MetaData ImagePostJSON        `form:"meta"`
}

type ImageUpdate struct {
	id int
}

type ImageDelete struct {
}
