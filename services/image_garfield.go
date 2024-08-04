package services

import (
	"bytes"
	"github.com/gofrs/uuid/v5"
	"github.com/joho/godotenv"
	"io"
	"log"
	"log/slog"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"webback/models"
	"webback/repositories"
	"webback/requests"
	"webback/responses"
)

type ImageGarfield struct {
	imageRepository repositories.Image
	fileRepository  repositories.File
	userRepository  repositories.User
	privatePath     string
	publicPath      string
	urlPath         string
}

func NewImageGarfield(
	imageRepository repositories.Image,
	userRepository repositories.User,
	fileRepository repositories.File,
) *ImageGarfield {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	privatePath := os.Getenv("PRIVATE_PATH")
	if privatePath == "" {
		log.Fatal("PRIVATE_PATH environment variable not set")
	}
	publicPath := os.Getenv("PUBLIC_PATH")
	if publicPath == "" {
		log.Fatal("PUBLIC_PATH environment variable not set")
	}
	urlPath := os.Getenv("URL_PATH")
	if urlPath == "" {
		log.Fatal("URL_PATH environment variable not set")
	}

	return &ImageGarfield{
		imageRepository,
		fileRepository,
		userRepository,
		privatePath,
		publicPath,
		urlPath,
	}
}

func (service *ImageGarfield) AddImage(request requests.ImagePost) (*responses.ImagePost, error) {

	// Get JSON from request
	imageJSON := request.MetaData

	// Check if public
	var basePath string
	if imageJSON.AllowPublic {
		basePath = service.publicPath
	} else {
		basePath = service.privatePath
	}

	// Generate file path and UUID
	filePath, fileUUID, err := service.generateFilePath(basePath)
	if err != nil {
		slog.Error("Error generating file path " + err.Error())
		return nil, err
	}
	filePath += "." + imageJSON.FileType

	// If public create url
	var urlPointer *string = nil
	if imageJSON.AllowPublic {
		url := service.urlPath + "/" + fileUUID + "." + imageJSON.FileType
		urlPointer = &url
	}

	// Open file
	fileHeader := request.FileData
	file, err := fileHeader.Open()
	if err != nil {
		slog.Error("Error opening file " + err.Error())
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file " + err.Error())
		}
	}(file)

	// Copy file to byte slice
	var bytesBuffer bytes.Buffer
	_, err = io.Copy(&bytesBuffer, file)
	if err != nil {
		return nil, err
	}
	byteSlice := bytesBuffer.Bytes()

	// Write byte slice to file storage
	err = service.fileRepository.Write(&byteSlice, filePath)
	if err != nil {
		return nil, err
	}

	// Create model
	image := models.Image{
		Path:        filePath,
		URL:         urlPointer,
		AllowPublic: imageJSON.AllowPublic,
		AllowML:     imageJSON.AllowML,
	}

	// Add image model to image repository
	err = service.imageRepository.Create(&image)
	if err != nil {
		println("Error creating image: " + err.Error())
		return nil, err
	}

	return &responses.ImagePost{Success: true}, nil
}

func (service *ImageGarfield) GetImage(id string) (*responses.ImageGet, error) {
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	imageModel, err := service.imageRepository.FindById(uint(uid))
	if err != nil {
		return nil, err
	}

	image := service.convertModelToResponse(*imageModel)
	return &image, nil
}

func (service *ImageGarfield) GetAllImagesPublic() (*[]responses.ImageGet, error) {
	var images []responses.ImageGet

	imageModels, err := service.imageRepository.FindAllWhereAllowPublic()
	if err != nil {
		return nil, err
	}

	for _, model := range *imageModels {
		images = append(images, service.convertModelToResponse(model))
	}

	return &images, nil
}

// Private functions

func (service *ImageGarfield) convertModelToResponse(image models.Image) responses.ImageGet {
	slog.Debug("Converting model to response")
	var displayName string
	if image.UserID == nil {
		displayName = "Anonymous"
	} else {
		user, err := service.userRepository.FindById(*image.UserID)
		if err != nil {
			displayName = "Anonymous"
		} else {
			displayName = user.DisplayName
		}
	}

	var tags []string
	if image.Tags != nil {
		strings.Fields(*image.Tags)
	} else {
		tags = []string{}
	}

	return responses.ImageGet{
		ID:         strconv.FormatUint(uint64(image.ID), 10),
		User:       displayName,
		UploadDate: image.CreatedAt.String(),
		URL:        image.Path,
		Tags:       tags,
	}
}

func (service *ImageGarfield) generateFilePath(basePath string) (string, string, error) {
	slog.Debug("Generating file path")
	fileUUID, err := uuid.NewV4()
	if err != nil {
		slog.Error("Error generating file uuid")
		return "", "", err
	}
	filePath := basePath + "/" + fileUUID.String()
	if service.fileRepository.FileExists(filePath) {
		return service.generateFilePath(basePath)
	}
	return filePath, fileUUID.String(), nil
}
