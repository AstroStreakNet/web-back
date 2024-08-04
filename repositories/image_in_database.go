package repositories

import (
	"errors"
	"gorm.io/gorm"
	"webback/models"
)

type ImageInDatabase struct {
	database *gorm.DB
}

// NewImageRepositoryInDatabase constructor for ImageInDatabase
func NewImageRepositoryInDatabase(database *gorm.DB) (*ImageInDatabase, error) {
	repository := &ImageInDatabase{database: database}
	err := repository.Initialize()
	if err != nil {
		return nil, err
	}
	return repository, nil
}

func (repository *ImageInDatabase) Initialize() error {
	return repository.database.AutoMigrate(&models.Image{})
}

func (repository *ImageInDatabase) FindAll() (*[]models.Image, error) {
	var images []models.Image
	repository.database.Find(&images)
	return &images, nil
}

func (repository *ImageInDatabase) FindById(id uint) (*models.Image, error) {
	var image models.Image
	result := repository.database.First(&image, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &ImageNotFound{}
		}
		return nil, result.Error
	}
	return &image, nil
}

func (repository *ImageInDatabase) FindAllWhereAllowPublic() (*[]models.Image, error) {
	var images []models.Image
	result := repository.database.Where(&models.Image{AllowPublic: true}).Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}
	return &images, nil
}

func (repository *ImageInDatabase) FindAllWhereAllowML() (*[]models.Image, error) {
	var images []models.Image
	result := repository.database.Where(&models.Image{AllowML: true}).Find(&images)
	if result.Error != nil {
		return nil, result.Error
	}
	return &images, nil
}

func (repository *ImageInDatabase) Create(image *models.Image) error {
	return repository.database.Create(image).Error
}

func (repository *ImageInDatabase) Update(image *models.Image) error {
	return repository.database.Save(image).Error
}

func (repository *ImageInDatabase) Delete(image *models.Image) error {
	result := repository.database.Delete(image)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &ImageNotFound{}
	}
	return nil
}
