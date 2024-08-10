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
	return repository.database.AutoMigrate(&models.ImageGorm{})
}

func (repository *ImageInDatabase) FindAll() (*[]models.Image, error) {
	var imagesGORM []models.ImageGorm
	var images []models.Image
	repository.database.Find(&imagesGORM)
	for _, imageGORM := range imagesGORM {
		images = append(images, models.ImageFromGorm(imageGORM))
	}
	return &images, nil
}

func (repository *ImageInDatabase) FindById(id uint) (*models.Image, error) {
	var imageGORM models.ImageGorm
	var image models.Image
	result := repository.database.First(&imageGORM, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &ImageNotFound{}
		}
		return nil, result.Error
	}
	image = models.ImageFromGorm(imageGORM)
	return &image, nil
}

func (repository *ImageInDatabase) FindAllWhereAllowPublic() (*[]models.Image, error) {
	var imagesGORM []models.ImageGorm
	var images []models.Image
	result := repository.database.Where(&models.ImageGorm{AllowPublic: true}).Find(&imagesGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, imageGORM := range imagesGORM {
		images = append(images, models.ImageFromGorm(imageGORM))
	}
	return &images, nil
}

func (repository *ImageInDatabase) FindAllWhereAllowML() (*[]models.Image, error) {
	var imagesGORM []models.ImageGorm
	var images []models.Image
	result := repository.database.Where(&models.ImageGorm{AllowML: true}).Find(&imagesGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	// Convert GORM to DTO
	for _, imageGORM := range imagesGORM {
		images = append(images, models.ImageFromGorm(imageGORM))
	}
	return &images, nil
}

func (repository *ImageInDatabase) Create(image *models.Image) error {
	imageGORM := models.GormFromImage(*image)
	return repository.database.Create(&imageGORM).Error
}

func (repository *ImageInDatabase) Update(image *models.Image) error {
	imageGORM := models.GormFromImage(*image)
	return repository.database.Save(&imageGORM).Error
}

func (repository *ImageInDatabase) Delete(image *models.Image) error {
	imageGORM := models.GormFromImage(*image)
	result := repository.database.Delete(&imageGORM)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &ImageNotFound{}
	}
	return nil
}

// private
