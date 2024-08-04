package repositories

type File interface {
	Initialize() error
	Read(path string) (*[]byte, error)
	Write(data *[]byte, path string) error
	Overwrite(data *[]byte, path string) error
	Delete(path string) error
	FileExists(path string) bool
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
