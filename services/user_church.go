package services

import (
	"errors"
	"net/mail"
	"webback/auth"
	"webback/models"
	"webback/repositories"
	"webback/requests"
)

type UserChurch struct {
	passwordEncoder auth.PasswordEncoder
	userRepository  repositories.User
}

func NewUserChurch(passwordEncoder auth.PasswordEncoder, userRepository repositories.User) *UserChurch {
	return &UserChurch{
		passwordEncoder: passwordEncoder,
		userRepository:  userRepository,
	}
}

func (service *UserChurch) Login(request *requests.Login) (*models.User, error) {
	if !service.isValidEmail(request.Email) {
		return nil, &InvalidEmail{}
	}

	user, err := service.userRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if service.passwordEncoder.CheckPassword(request.Password, user.Password) {
		return user, nil
	}
	return nil, &InvalidPassword{}
}

func (service *UserChurch) Register(request *requests.Register) error {
	if !service.isValidEmail(request.Email) {
		return &InvalidEmail{}
	}

	_, err := service.userRepository.FindByEmail(request.Email)
	// If user with email exists already
	if !errors.Is(err, &repositories.UserNotFound{}) {
		return &EmailAlreadyInUse{}
	}

	hashedPassword, err := service.passwordEncoder.EncodePassword(request.Password)
	if err != nil {
		return err
	}

	userBuilder := models.NewUserBuilder().
		WithEmail(request.Email).
		WithPassword(hashedPassword).
		WithDisplayName(request.DisplayName)

	if request.FirstName != "" {
		userBuilder.WithFirstName(request.FirstName)
	}
	if request.LastName != "" {
		userBuilder.WithLastName(request.LastName)
	}

	user := userBuilder.Build()
	return service.userRepository.Create(user)
}

// Private

func (service *UserChurch) isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}
