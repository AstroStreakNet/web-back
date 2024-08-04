package repositories

import (
	"bytes"
	"errors"
	"io"
	"os"
)

type FileOnSystem struct {
}

func NewFileRepositoryOnSystem() *FileOnSystem {
	return &FileOnSystem{}
}

func (repository *FileOnSystem) Initialize() error {
	return nil
}

func (repository *FileOnSystem) Read(path string) (*[]byte, error) {
	if !repository.FileExists(path) {
		return nil, &FileDoesNotExist{}
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// TODO logging
		}
	}(file)

	var dataBuffer bytes.Buffer
	_, err = io.Copy(&dataBuffer, file)
	data := dataBuffer.Bytes()
	return &data, nil
}

func (repository *FileOnSystem) Write(data *[]byte, path string) error {
	if repository.FileExists(path) {
		return &FileAlreadyExists{}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// TODO logging
		}
	}(file)

	_, err = io.Copy(file, bytes.NewReader(*data))
	if err != nil {
		return err
	}

	return nil
}

func (repository *FileOnSystem) Overwrite(data *[]byte, path string) error {
	if !repository.FileExists(path) {
		return &FileDoesNotExist{}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// TODO logging
		}
	}(file)

	_, err = io.Copy(file, bytes.NewReader(*data))
	if err != nil {
		return err
	}

	return nil
}

func (repository *FileOnSystem) Delete(path string) error {
	if !repository.FileExists(path) {
		return &FileDoesNotExist{}
	}
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func (repository *FileOnSystem) FileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
