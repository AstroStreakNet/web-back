package main

import (
	"fmt"
	"github.com/AstroStreakNet/telescope/astrometry"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"webback/controllers"
	"webback/repositories"
	"webback/services"
	"webback/setup"
)

func main() {

	// Get environment variables
	privatePath := os.Getenv("PRIVATE_PATH")
	if privatePath == "" {
		log.Fatal("PRIVATE_PATH environment variable not set")
	}
	publicPath := os.Getenv("PUBLIC_PATH")
	if publicPath == "" {
		log.Fatal("PUBLIC_PATH environment variable not set")
	}
	urlPath := os.Getenv("URL_PATH")
	if urlPath == "" {
		log.Fatal("URL_PATH environment variable not set")
	}

	// Connect to database
	database := setup.NewMySQLDatabaseConnection()

	// Repositories
	imageRepository, err := repositories.NewImageRepositoryInDatabase(database)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create image repository: %v", err))
	}

	userRepository, err := repositories.NewUserRepositoryInDatabase(database)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create user repository: %v", err))
	}

	fileRepository := repositories.NewFileRepositoryOnSystem(privatePath, publicPath)

	// Proxies
	astrometryProxy := astrometry.NewAstrometryClient("placeholder")

	// Services
	imageService := services.NewImageGarfield(
		imageRepository, userRepository, fileRepository, *astrometryProxy, &urlPath)

	authService := services.NewAuthJonesy(
		userRepository, 30, "placeholder")

	// Controllers
	imageController := controllers.NewImageController(imageService, authService)

	// Setup router
	router := gin.Default()

	// Setup route groups
	// admin := router.Group("/admin")
	// user := router.Group("/user")
	image := router.Group("/image")

	// Setup static image serving
	router.Static(urlPath, publicPath)

	// Assign routes to controller methods
	image.GET("", imageController.GetImage)
	image.POST("", imageController.PostImage)

	// Run on port 8080
	err = router.Run(":8090")
	if err != nil {
		return
	}
}
