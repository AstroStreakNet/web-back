package web_back

import (
	"github.com/gin-gonic/gin"
	"webback/controllers"
	"webback/models"
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

	// Run
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
