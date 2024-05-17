// controllers/image.go

package controllers

import (
	// "encoding/json"
	"fmt"
	// "log"
	"mime/multipart"
	"net/http"
	// "webback/models"

	"github.com/gin-gonic/gin"
)

// Request Structure
type PostImageRequest struct {
	Image       *multipart.FileHeader `form:"image"`
	AllowPublic bool                  `json:"allowPublic"`
	AllowML     bool                  `json:"allowML"`
}

type ImageCriteria struct {
    Name        string              `form:"name"`
    Uploader    string              `form:"uploader"`
    UploadDate  int                 `form:"uploadDate"`
    URL         string              `form:"url"`
    TAGS      []string              `form:"tags"`
}

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

