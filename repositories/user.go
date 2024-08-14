package repositories

import "webback/models"

type User interface {
	Initialize() error
	FindAll() (*[]models.User, error)
	FindById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByFirstName(name string) (*[]models.User, error)
	FindByLastName(name string) (*[]models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}

// Errors

type UserNotFound struct{}

func (e *UserNotFound) Error() string {
	return "user not found"
}
