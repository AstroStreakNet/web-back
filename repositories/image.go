package repositories

import "webback/models"

type Image interface {
	Initialize() error
	FindAll() (*[]models.Image, error)
	FindById(id uint) (*models.Image, error)
	FindAllWhereAllowPublic() (*[]models.Image, error)
	FindAllWhereAllowML() (*[]models.Image, error)
	Create(*models.Image) error
	Update(*models.Image) error
	Delete(*models.Image) error
}

// Errors

type ImageNotFound struct{}

func (e *ImageNotFound) Error() string {
	return "image not found"
}
