package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"webback/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ImageCriteria struct {
	Name       string   `form:"name"`
	Uploader   string   `form:"uploader"`
	UploadDate int      `form:"uploadDate"`
	URL        string   `form:"url"`
	TAGS       []string `form:"tags"`
}

// Request Structure
type UploadTags struct {
	Image       *multipart.FileHeader `form:"image"`
	AllowPublic bool                  `json:"allowPublic"`
	AllowML     bool                  `json:"allowML"`
	Telescope   string                `json:"telescope"`
	Observatory string                `json:"observatory"`
	RightAscen  string                `json:"rightAscen"`
	Declination string                `json:"declination"`
	Julian      time.Time             `json:"julian"`
	Exposure    time.Time             `json:"exposure"`
	StreakType  string                `json:"streakType"`
	Creation    string                `json:"creationDate"`
	// tags string array
	// cords
}

// PostImage handles HTTP POST requests for uploading images
func PostImage(c *gin.Context) {
	var req UploadTags
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if an image is uploaded
	if req.Image == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image uploaded"})
		return
	}

	// Get .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	mediaPath := os.Getenv("MEDIA_PATH")

	// Save the uploaded image to a specific location
	imagePath := mediaPath + req.Image.Filename
	if err := c.SaveUploadedFile(req.Image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	processImage(imagePath, req.AllowPublic, req.AllowML, req.Telescope,
		req.Observatory, req.RightAscen, req.Declination, req.Julian, req.Exposure)

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

// Process image and add to database
func processImage(imagePath string, allowPublic bool, allowML bool, telescope string,
	observatory string, rightAscen string, declination string, date time.Time,
	exposure time.Time) {

	// call telescope with path to image

	// telescope will return cords and contains which is required for the database
	// add image to database
}

// GetImageAll handles HTTP GET requests to retrieve all public images from the database
func GetImageAll(c *gin.Context) {
	var criteria = ImageCriteria{}

	if err := c.BindJSON(&criteria); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// process image criterias

	// only for testing
	fmt.Printf("[request] Name=%v Uploader=%v UDate=%v URL=%v TAGS=%v",
		criteria.Name, criteria.Uploader, criteria.UploadDate, criteria.URL, criteria.TAGS)

	c.JSON(http.StatusOK, gin.H{"message": "request received"})
}

func CreateImage(c *gin.Context) {
	var newImage models.Image
	if err := c.ShouldBindJSON(&newImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	image := models.CreateImage(&newImage)
	c.JSON(http.StatusOK, image)
}

func DeleteImage(c *gin.Context) {
	imageID := c.Param("imageID")
	err := models.DeleteImage(imageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}

func UpdateImage(c *gin.Context) {
	var updateImage models.Image
	if err := c.ShouldBindJSON(&updateImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	imageID := c.Param("imageID")
	imageDetails, err := models.GetImageByID(imageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Image not found"})
		return
	}

	if updateImage.ImagePath != "" {
		imageDetails.ImagePath = updateImage.ImagePath
	}
	if updateImage.ObservatoryCode != "" {
		imageDetails.ObservatoryCode = updateImage.ObservatoryCode
	}
	if !updateImage.TimeOfObservation.IsZero() {
		imageDetails.TimeOfObservation = updateImage.TimeOfObservation
	}
	if updateImage.StreakType != "" {
		imageDetails.StreakType = updateImage.StreakType
	}
	if updateImage.UserID != "" {
		imageDetails.UserID = updateImage.UserID
	}
	if updateImage.AllowPublic != imageDetails.AllowPublic {
		imageDetails.AllowPublic = updateImage.AllowPublic
	}
	if updateImage.AllowML != imageDetails.AllowML {
		imageDetails.AllowML = updateImage.AllowML
	}

	// Save the updated image
	if err := imageDetails.Update(DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image"})
		return
	}

	c.JSON(http.StatusOK, imageDetails)
}
