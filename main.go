package main

import (
	"webback/controllers"
	"webback/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database connection
	models.Connect()

	// Setup router
	r := gin.Default()

	// Setup static image serving
	r.Static("/public", "./static")

	// Add controllers
	r.GET("/image", controllers.GetImageAll)
	r.POST("/CLI", controllers.HandleBrowseRequest)
	r.POST("/upload", controllers.PostImage)

	// Run on port 8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
