package models

import (
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(nil, nil) // TODO: setup db login and config
	if err != nil {
		log.Fatal(err)
	}

	// Migrations

	DB = database
}
