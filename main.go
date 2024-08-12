package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/AstroStreakNet/telescope/astrometry"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"webback/controllers"
	"webback/repositories"
	"webback/services"
)

// Database connections

func newInMemoryDatabaseConnection() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return database
}

func newMySQLDatabaseConnection() *gorm.DB {
	database, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return database
}

func newFirestoreConnection(ctx context.Context, projectID, credentialsPath string) *firestore.Client {
	clientOptions := option.WithCredentialsFile(credentialsPath)
	client, err := firestore.NewClient(ctx, projectID, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {

	// Load .env
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	// Get .env variables
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
	database := newInMemoryDatabaseConnection()

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
		imageRepository, userRepository, fileRepository, *astrometryProxy, urlPath)

	// Controllers
	imageController := controllers.NewImageController(imageService)

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
