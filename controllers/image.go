package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"strings"
	"webback/requests"
	"webback/responses"
	"webback/services"
)

type Image struct {
	imageService services.Image
	authService  services.Auth
}

// NewImageController Constructor for Image
func NewImageController(imageService services.Image, authService services.Auth) *Image {
	return &Image{
		imageService,
		authService,
	}
}

func (controller *Image) GetImage(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		response, err := controller.imageService.GetImage(id)
		if err != nil {
			// TODO replace err.Error() with ambiguous message
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}

	response, err := controller.imageService.GetAllImagesPublic()
	if err != nil {
		// TODO replace err.Error() with ambiguous message
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller *Image) PostImage(c *gin.Context) {

	var request requests.ImagePost
	if err := c.ShouldBind(&request); err != nil {
		// TODO replace err.Error() with ambiguous message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check JSON
	json := request.MetaData
	if !controller.validFileType(json.FileType) || json.FileType == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid file type"})
		return
	}

	user := c.Query("user")
	if user != "" {
		requestHeader := c.Request.Header.Get("Authorization")
		token, claims, err := controller.authService.ParseToken(requestHeader)
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized",
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		if claims.User != user {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth for requested user"})
			return
		}
	}
	err := controller.imageService.AddImage(request, user)

	if err != nil {
		if errors.Is(err, services.IncidentalError{}) {
			c.JSON(http.StatusOK, &responses.ImagePost{Message: err.Error()})
			return
		} else {
			// TODO replace err.Error() with ambiguous message
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, &responses.ImagePost{Message: "image successfully uploaded"})
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
