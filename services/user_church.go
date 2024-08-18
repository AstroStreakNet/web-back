package services

import (
	"errors"
	"net/mail"
	"strconv"
	"webback/auth"
	"webback/models"
	"webback/repositories"
	"webback/requests"
	"webback/responses"
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

func (service *UserChurch) GetUser(email string) (*models.User, error) {
	return service.userRepository.FindByEmail(email)
}

func (service *UserChurch) GetDetails(userID string) (*responses.GetDetails, error) {

	userIDUint64, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return nil, err
	}

	userIDUint := uint(userIDUint64)

	user, err := service.userRepository.FindById(userIDUint)
	if err != nil {
		return nil, err
	}

	var firstName, lastName string
	if user.FirstName == nil {
		firstName = ""
	} else {
		firstName = *user.FirstName
	}
	if user.LastName == nil {
		lastName = ""
	} else {
		lastName = *user.LastName
	}

	return &responses.GetDetails{
		Email:       user.Email,
		FirstName:   firstName,
		LastName:    lastName,
		DisplayName: user.DisplayName,
	}, nil
}

func (service *UserChurch) UpdateDetails() {

}

// Private

func (service *UserChurch) isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}
