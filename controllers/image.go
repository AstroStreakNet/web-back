// controllers/image.go

package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"webback/models"
)

// PostImage handles HTTP POST requests for uploading images
func PostImage(c *gin.Context) {
	var req PostImageRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// image := req.Image
}

// GetImageAll handles HTTP GET requests to retrieve all public images from the database
func GetImageAll(c *gin.Context) {
	// Query database for public images
	var images []models.Image
	models.DB.Table("images").Where("allow_public = ?", true).Find(&images)

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

// Request Structure

type PostImageRequest struct {
	Image       *multipart.FileHeader `form:"image"`
	AllowPublic bool                  `json:"allowPublic"`
	AllowML     bool                  `json:"allowML"`
}

