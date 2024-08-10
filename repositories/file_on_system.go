package repositories

import (
	"bytes"
	"errors"
	"github.com/gofrs/uuid/v5"
	"image/jpeg"
	"image/png"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"webback/FITS"
)

type FileOnSystem struct {
	privatePath string
	publicPath  string
}

func NewFileRepositoryOnSystem(privatePath, publicPath string) *FileOnSystem {
	return &FileOnSystem{
		privatePath,
		publicPath,
	}
}

func (repository *FileOnSystem) Initialize() error {
	return nil
}

func (repository *FileOnSystem) Read(fileName string) (*bytes.Buffer, error) {
	if !repository.FileExists(fileName) {
		return nil, &FileDoesNotExist{}
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// TODO logging
		}
	}(file)

	var data bytes.Buffer
	_, err = io.Copy(&data, file)
	return &data, nil
}

func (repository *FileOnSystem) Write(data *bytes.Buffer, fileName string) error {
	if repository.FileExists(fileName) {
		return &FileAlreadyExists{}
	}

	filePath := repository.privatePath + "/" + fileName

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			slog.Error("Error while closing file", slog.String("path", filePath))
		}
	}(file)

	_, err = io.Copy(file, data)
	if err != nil {
		return err
	}

	return nil
}

func (repository *FileOnSystem) WritePreview(data *bytes.Buffer, fileName string) error {
	// Get preview file name
	fileType := filepath.Ext(fileName)
	previewFileName := strings.TrimSuffix(fileName, fileType) + ".jpeg"

	if repository.PreviewExists(previewFileName) {
		return &FileAlreadyExists{}
	}
	filePath := repository.publicPath + "/" + previewFileName

	// Convert image data to jpeg
	preview, err := repository.convertImageToJPEG(data, fileType)
	if err != nil {
		return err
	}

	// Create file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			slog.Error("Error while closing file", slog.String("path", filePath))
		}
	}(file)

	// Copy data to file
	_, err = io.Copy(file, preview)
	if err != nil {
		return err
	}

	return nil
}

func (repository *FileOnSystem) Overwrite(data *bytes.Buffer, fileName string) error {
	if !repository.FileExists(fileName) {
		return &FileDoesNotExist{}
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// TODO logging
		}
	}(file)

	_, err = io.Copy(file, data)
	if err != nil {
		return err
	}

	return nil
}

func (repository *FileOnSystem) Delete(fileName string) error {
	if !repository.FileExists(fileName) {
		return &FileDoesNotExist{}
	}
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func (repository *FileOnSystem) DeletePreview(fileName string) error {
	if !repository.FileExists(fileName) {
		return &FileDoesNotExist{}
	}
	err := os.Remove(repository.publicPath + "/" + fileName)
	if err != nil {
		return err
	}
	return nil
}

func (repository *FileOnSystem) GenerateFileName(fileType string) (string, error) {
	fileUUID, err := uuid.NewV4()
	if err != nil {
		slog.Error("Error generating file uuid")
		return "", err
	}
	fileName := fileUUID.String() + "." + fileType

	if repository.FileExists(repository.privatePath + "/" + fileName) {
		return repository.GenerateFileName(fileType)
	}
	return fileName, nil
}

func (repository *FileOnSystem) FileExists(fileName string) bool {
	if _, err := os.Stat(repository.privatePath + "/" + fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (repository *FileOnSystem) PreviewExists(fileName string) bool {
	if _, err := os.Stat(repository.publicPath + "/" + fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (repository *FileOnSystem) convertImageToJPEG(data *bytes.Buffer, currentType string) (*bytes.Buffer, error) {
	var newData bytes.Buffer
	switch currentType {

	case ".jpeg", ".jpg":
		return data, nil

	case ".png":
		image, err := png.Decode(data)
		if err != nil {
			return nil, err
		}
		err = jpeg.Encode(&newData, image, &jpeg.Options{Quality: 75})
		if err != nil {
			return nil, err
		}
		return &newData, nil

	case ".fits":
		image, err := FITS.Decode(data)
		if err != nil {
			return nil, err
		}
		err = jpeg.Encode(&newData, image, &jpeg.Options{Quality: 75})
		if err != nil {
			return nil, err
		}
		return &newData, nil

	default:
		return nil, errors.New("unsupported file type")
	}
}
