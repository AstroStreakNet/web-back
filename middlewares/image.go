package middlewares

import (
	"github.com/AstroStreakNet/telescope/astrometry"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

var astrometryClient *astrometry.Client

func getAstrometryClient() *astrometry.Client {
	if astrometryClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if astrometryClient == nil {
			astrometryClient = astrometry.NewAstrometryClient("")
		}
	}
	return astrometryClient
}

// ProcessImage Process image and add to database
func ProcessImage(imagePath string, allowPublic bool, allowML bool, telescope string,
	observatory string, rightAscen string, declination string, date string,
	exposure string) {

	// Get client instance
	client := getAstrometryClient()

	// Call telescope with path to image
	if _, err := client.Connect(); err != nil {
		log.Fatal(err)
	}

	subID, err := client.UploadFile(imagePath)
	if err != nil {
		return
	}

	// add image to database
}
