package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	imagePath string
	fileType  string // May or may not be needed
}
