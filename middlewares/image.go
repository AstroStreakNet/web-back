package middlewares

import (
	"github.com/AstroStreakNet/telescope/astrometry"
	"gorm.io/datatypes"
	"log"
	"strconv"
	"sync"
	"time"
	"webback/models"
)

var lock = &sync.Mutex{}

var astrometryClient *astrometry.Client

func getAstrometryClient() *astrometry.Client {
	if astrometryClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if astrometryClient == nil {
			astrometryClient = astrometry.NewAstrometryClient("") // TODO: load apiKey through .env or something
		}
	}
	return astrometryClient
}

// AddImage add image to database
func AddImage(allowPublic, allowML bool, imagePath, observatory, rightAscension, declination, julianDate,
	exposureTime string) {

	// Convert datatypes
	var floatRA float32
	value, err := strconv.ParseFloat(rightAscension, 32)
	if err != nil {
		log.Println(err)
	} else {
		floatRA = float32(value)
	}

	var floatDEC float32
	value, err = strconv.ParseFloat(declination, 32)
	if err != nil {
		log.Println(err)
	} else {
		floatDEC = float32(value)
	}

	var dateJD time.Time
	dateJD, err = time.Parse("", julianDate) // TODO: add datetime format/layout
	if err != nil {
		log.Println(err)
	}

	var timeED datatypes.Time
	err = timeED.Scan(exposureTime)
	if err != nil {
		log.Println(err)
	}

	// Create image model
	img := models.Image{
		ImagePath:        imagePath,
		AllowPublic:      allowPublic,
		AllowML:          allowML,
		ObservatoryCode:  observatory,
		RightAscension:   floatRA,
		Declination:      floatDEC,
		JulianDate:       dateJD,
		ExposureDuration: timeED,
		AstroSubID:       0,
		AstroProcessed:   false,
	}

	models.DB.
		Table("images").
		Create(&img)
}

// ProcessImage Begin image processing, return submission id
func ProcessImage(imagePath string) int {

	// Get client instance
	client := getAstrometryClient()

	// Connect client to astrometry, needed for uploads
	if _, err := client.Connect(); err != nil {
		log.Fatal(err)
	}

	// Call telescope with path to image
	subID, err := client.UploadFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	return subID
}

// CheckImages go through all images and their submission IDs to see if they have been processed
func CheckImages() {

	// Get client instance
	client := getAstrometryClient()

	// Query database for images
	var images []models.Image
	models.DB.
		Table("images").
		Where("astro_processed = ? AND astro_sub_id <> 0", []string{"0", "0"}).
		Find(&images)

	// Loop through submissions ids

	for _, image := range images {
		subID := image.AstroSubID

		partialReview, err := client.GetPartialReview(subID)
		if err != nil {
			log.Fatal(err)
		}

		if partialReview.Finished {
			image.AstroProcessed = partialReview.Finished
			// TODO: add other stuff
		} else {
			continue
		}
	}

	// TODO: update database entries with new data
}
