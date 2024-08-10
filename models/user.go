// models/user.go

package models

import "time"

type User struct {
	ID          uint
	CreatedAt   time.Time
	Email       string
	Password    string
	DisplayName string
	FirstName   string
	LastName    string
	Role        string
	Images      []uint
}
