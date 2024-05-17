// controllers/image.go

package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	// "encoding/json"
	// "webback/models"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

type ImageCriteria struct {
    Name        string               `form:"name"`
    Uploader    string               `form:"uploader"`
    UploadDate  int                  `form:"uploadDate"`
    URL         string               `form:"url"`
    TAGS      []string               `form:"tags"`
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
    Julian      string                `json:"julian"`
    Exposure    string                `json:"exposure"`
    StreakType  string                `json:"streakType"`
}

// PostImage handles HTTP POST requests for uploading images
func PostImage(c *gin.Context) {
	var req UploadTags
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	imagePath := (mediaPath + req.Image.Filename)
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
    observatory string, rightAscen string, declination string, date string, 
    exposure string) {

    // call telescope with path to image 
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

