package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"webback/controllers"
	"webback/repositories"
	"webback/services"
)

func newDatabaseConnection() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return database
}

func main() {
	// Connect to database
	database := newDatabaseConnection()

	// Instantiate repositories
	imageRepository, err := repositories.NewImageRepositoryInDatabase(database)
	if err != nil {
		panic(fmt.Sprintf("failed to create image repository: %v", err))
	}
	userRepository, err := repositories.NewUserRepositoryInDatabase(database)
	if err != nil {
		panic(fmt.Sprintf("failed to create user repository: %v", err))
	}
	fileRepository := repositories.NewFileRepositoryOnSystem()

	// Instantiate services
	imageService := services.NewImageGarfield(imageRepository, userRepository, fileRepository)

	// Instantiate controllers
	imageController := controllers.NewImageController(imageService)

	// Setup router
	router := gin.Default()

	// Setup route groups
	// admin := router.Group("/admin")
	// user := router.Group("/user")
	image := router.Group("/image")

	// Setup static image serving
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	urlPath := os.Getenv("URL_PATH")
	if urlPath == "" {
		log.Fatal("URL_PATH environment variable not set")
	}
	publicPath := os.Getenv("PUBLIC_PATH")
	if publicPath == "" {
		log.Fatal("PUBLIC_PATH environment variable not set")
	}
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
