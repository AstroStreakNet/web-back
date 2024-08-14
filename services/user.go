package services

import (
	"webback/models"
	"webback/requests"
)

type User interface {
	Login(request *requests.Login) (*models.User, error)
	Register(request *requests.Register) error
	GetUser(email string) (*models.User, error)
	GetDetails()
	UpdateDetails()
}

// Errors
type InvalidEmail struct{}

func (e InvalidEmail) Error() string {
	return "invalid email address"
}

type InvalidPassword struct{}

func (e InvalidPassword) Error() string {
	return "invalid password"
}

type EmailAlreadyInUse struct{}

func (e EmailAlreadyInUse) Error() string {
	return "email already in use"
}
