package web_back

import (
	"github.com/gin-gonic/gin"
	"webback/models"
)

func main() {
	// Setup router
	r := gin.Default()

	// Database connection
	models.Connect()

	// Run
	err := r.Run()
	if err != nil {
		return
	}
}
