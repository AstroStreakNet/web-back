package controllers

import (
	"github.com/gin-gonic/gin"
	"webback/models"
)

func UploadImage(c *gin.Context) {

}

func GetImages(c *gin.Context) {
	var users []models.User
	models.DB.Where("allow_public = ?", true).Find(&users)
}
