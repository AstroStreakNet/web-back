package models

import "gorm.io/datatypes"

type gender string

const (
	MALE      gender = "MALE"
	FEMALE    gender = "FEMALE"
	UNDEFINED gender = "UNDEFINED"
)

type User struct {
	ID           string         `gorm:"primaryKey; column:user_id; type:VARCHAR(45)"`
	AstronomerID uint           `gorm:"column:astronomer_id; type: INT"`
	Password     string         `gorm:"column:#password; type:VARCHAR(45)"`
	FirstName    string         `gorm:"column:first_name; type:VARCHAR(45)"`
	LastName     string         `gorm:"column:last_name; type:VARCHAR(45)"`
	DOB          datatypes.Date `gorm:"column:dob; type:DATE"`
	Gender       gender         `gorm:"column:gender; type:ENUM('MALE', 'FEMALE', 'UNDEFINED')"`
	Permissions  uint           `gorm:"column:permissions; type:TINYINT"`
	ImagePublish uint           `gorm:"column:image_publish; type:TINYINT"`
}
