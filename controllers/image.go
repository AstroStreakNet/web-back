package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webback/models"
)

// Handle HTTP POST requests for uploading images
func PostImage(c *gin.Context) {

}

// Handles HTTP GET requests to retrieve all public images from the database
func GetImageAll(c *gin.Context) {
	// Query database for public images
	var images []models.Image
	models.DB.Where("allow_public = ?", true).Find(&images)

	// Marshall data to json
	jsonData, err := json.Marshal(images)
	if err != nil {
		// Abort if failure
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatalf("Failed to marshall images: %s", err)
	}

	// Return data with 200 response
	c.JSON(http.StatusOK, jsonData)
}

