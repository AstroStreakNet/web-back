package repositories

import "bytes"

type File interface {
	Initialize() error
	Read(path string) (*bytes.Buffer, error)
	Write(data *bytes.Buffer, fileName string) error
	WritePreview(data *bytes.Buffer, fileName string) (*string, error)
	Overwrite(data *bytes.Buffer, fileName string) error
	Delete(fileName string) error
	DeletePreview(fileName string) error
	GenerateFileName(fileType string) (string, error)
	FileExists(fileName string) bool
	PreviewExists(fileName string) bool
	GetFilePath(fileName string) string
	GetPreviewPath(fileName string) string
}

// Errors

type FileAlreadyExists struct{}

func (e *FileAlreadyExists) Error() string {
	return "file already exists"
}

type FileDoesNotExist struct{}

func (e *FileDoesNotExist) Error() string {
	return "file does not exist"
}
