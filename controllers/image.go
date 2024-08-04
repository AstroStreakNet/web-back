package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webback/requests"
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
	response, err := controller.imageService.AddImage(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
