package services

import (
	"bytes"
	"github.com/AstroStreakNet/telescope/astrometry"
	"github.com/gofrs/uuid/v5"
	"io"
	"log/slog"
	"mime/multipart"
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
	astrometryProxy astrometry.Client
	urlPath         string
}

func NewImageGarfield(
	imageRepository repositories.Image,
	userRepository repositories.User,
	fileRepository repositories.File,
	astrometryProxy astrometry.Client,
	urlPath string,
) *ImageGarfield {

	return &ImageGarfield{
		imageRepository,
		fileRepository,
		userRepository,
		astrometryProxy,
		urlPath,
	}
}

func (service *ImageGarfield) AddImage(request requests.ImagePost) (*responses.ImagePost, error) {

	// Begin Image creation
	imageBuilder := models.NewImageBuilder()

	// Get JSON from request
	imageJSON := request.MetaData
	imageBuilder.WithAllowPublic(imageJSON.AllowPublic)
	imageBuilder.WithAllowML(imageJSON.AllowML)

	// Generate file name
	fileName, err := service.fileRepository.GenerateFileName(imageJSON.FileType)
	if err != nil {
		slog.Error("Error generating file path " + err.Error())
		return nil, err
	}
	imageBuilder.WithPath(fileName)

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

	// Copy file to byteBuffers
	var bytesBuffer bytes.Buffer

	// If public create preview image
	if imageJSON.AllowPublic {
		var bytesBufferClone bytes.Buffer
		tee := io.TeeReader(file, &bytesBufferClone)
		_, err = io.Copy(&bytesBuffer, tee)
		if err != nil {
			return nil, err
		}

		err = service.fileRepository.Write(&bytesBuffer, fileName)
		if err != nil {
			return nil, err
		}
		err = service.fileRepository.WritePreview(&bytesBufferClone, fileName)
		if err != nil {
			slog.Error("Error writing preview file " + err.Error())
			return nil, err
		}

		imageBuilder.WithURL(service.urlPath + "/" + fileName)

	} else {
		_, err = io.Copy(&bytesBuffer, file)
		if err != nil {
			return nil, err
		}
		err = service.fileRepository.Write(&bytesBuffer, fileName)
		if err != nil {
			return nil, err
		}

	}

	// Build image
	image := imageBuilder.Build()

	// Add image model to image repository
	err = service.imageRepository.Create(image)
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
	if image.UserID == 0 {
		displayName = "Anonymous"
	} else {
		user, err := service.userRepository.FindById(image.UserID)
		if err != nil {
			displayName = "Anonymous"
		} else {
			displayName = user.DisplayName
		}
	}

	var tags []string
	if image.Tags != "" {
		strings.Fields(image.Tags)
	} else {
		tags = []string{}
	}

	return responses.ImageGet{
		ID:         strconv.FormatUint(uint64(image.ID), 10),
		User:       displayName,
		UploadDate: image.CreatedAt.String(),
		URL:        image.URL,
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
