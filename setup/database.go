package setup

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"google.golang.org/api/option"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func NewInMemoryDatabaseConnection() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return database
}

type DatabaseUserInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewMySQLDatabaseConnection() *gorm.DB {
	mysqlHost := os.Getenv("MYSQL_HOST")
	// mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	// Get user info
	filePath, err := filepath.Abs(os.Getenv("MYSQL_SECRET_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var userInfo DatabaseUserInfo
	err = json.Unmarshal(info, &userInfo)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	// Create DSN from variables
	dsn := userInfo.User + ":" + userInfo.Password + "@tcp(" + mysqlHost + ")/" + mysqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Create connection
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database at " + mysqlHost)
	}
	return database
}

func NewFirestoreConnection(ctx context.Context, projectID, credentialsPath string) *firestore.Client {
	clientOptions := option.WithCredentialsFile(credentialsPath)
	client, err := firestore.NewClient(ctx, projectID, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
