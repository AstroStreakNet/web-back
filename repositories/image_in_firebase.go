package repositories

import (
	"cloud.google.com/go/firestore"
	"webback/models"
)

type ImageInFirebase struct {
	client *firestore.Client
}

func NewImageRepositoryInFirebase() *ImageInFirebase {
	return &ImageInFirebase{}
}

func (repository *ImageInFirebase) Initialize() error {
	return nil
}

func (repository *ImageInFirebase) FindAll() (*[]models.Image, error) {
	return nil, nil
}

func (repository *ImageInFirebase) FindById(id uint) (*models.Image, error) {
	return nil, nil
}

func (repository *ImageInFirebase) FindAllWhereAllowPublic() (*[]models.Image, error) {
	return nil, nil
}

func (repository *ImageInFirebase) FindAllWhereAllowML() (*[]models.Image, error) {
	return nil, nil
}

func (repository *ImageInFirebase) Create(*models.Image) error {
	return nil
}

func (repository *ImageInFirebase) Update(*models.Image) error {
	return nil
}

func (repository *ImageInFirebase) Delete(*models.Image) error {
	return nil
}
