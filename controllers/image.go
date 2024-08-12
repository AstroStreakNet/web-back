package controllers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
	"webback/requests"
	"webback/responses"
	"webback/services"
)

type Image struct {
	imageService services.Image
}

// NewImageController Constructor for Image
func NewImageController(imageService services.Image) *Image {
	return &Image{
		imageService,
	}
}

func (controller *Image) GetImage(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		response, err := controller.imageService.GetImage(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}

	response, err := controller.imageService.GetAllImagesPublic()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller *Image) PostImage(c *gin.Context) {

	var request requests.ImagePost
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check JSON
	json := request.MetaData
	if !controller.validFileType(json.FileType) || json.FileType == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid file type"})
		return
	}

	var response *responses.ImagePost
	var err error

	user := c.Query("user")
	if user != "" {

	} else {
		response, err = controller.imageService.AddImage(request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, response)
}

func (controller *Image) validFileType(fileType string) bool {
	fileTypeLower := strings.ToLower(fileType)
	slog.Debug(fileTypeLower)
	validTypes := []string{"jpeg", "jpg", "png", "fits"}
	for _, validType := range validTypes {
		if fileTypeLower == validType {
			return true
		}
	}
	return false
}
