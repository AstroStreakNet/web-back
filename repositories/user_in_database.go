package repositories

import (
	"errors"
	"gorm.io/gorm"
	"webback/models"
)

type UserInDatabase struct {
	database *gorm.DB
}

func NewUserRepositoryInDatabase(database *gorm.DB) (*UserInDatabase, error) {
	repository := &UserInDatabase{database: database}
	err := repository.Initialize()
	if err != nil {
		return nil, err
	}
	return repository, nil
}

func (repository *UserInDatabase) Initialize() error {
	return repository.database.AutoMigrate(&models.User{})
}

func (repository *UserInDatabase) FindAll() (*[]models.User, error) {
	var users []models.User
	result := repository.database.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (repository *UserInDatabase) FindById(id uint) (*models.User, error) {
	var user *models.User
	result := repository.database.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &UserNotFound{}
		}
		return nil, result.Error
	}
	return user, nil
}

func (repository *UserInDatabase) FindByFirstName(name string) (*[]models.User, error) {
	var users []models.User
	result := repository.database.Where("first_name = ?", name).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (repository *UserInDatabase) FindByLastName(name string) (*[]models.User, error) {
	var users []models.User
	result := repository.database.Where("last_name = ?", name).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func (repository *UserInDatabase) Create(user *models.User) error {
	return repository.database.Create(user).Error
}

func (repository *UserInDatabase) Update(user *models.User) error {
	// As long as the model has a valid primary key this will UPDATE the appropriate user.
	// If the primary key of the model can't be found this will CREATE the user instead.
	return repository.database.Save(user).Error
}

func (repository *UserInDatabase) Delete(user *models.User) error {
	result := repository.database.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &UserNotFound{}
	}
	return nil
}
