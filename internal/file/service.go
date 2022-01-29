package file

import (
	"fmt"
	"io/ioutil"
)

// Service defines the available functions that a file.service provides.
type Service interface {
	GetFile(name string) ([]byte, bool)
	StoreFile(name string, file []byte) error
}

type service struct{}

// NewFileService returns an instance of file.service.
func NewFileService() *service {
	return &service{}
}

// GetFile returns the file based on the given name, if the file does not exist it returns nil, false.
func (s *service) GetFile(name string) ([]byte, bool) {
	file, err := ioutil.ReadFile(fmt.Sprintf("uploads/%s", name))

	return file, err == nil
}

// StoreFile stores file with the given name to the filesystem under `uploads` folder.
func (s *service) StoreFile(name string, file []byte) error {
	return ioutil.WriteFile(fmt.Sprintf("uploads/%s", name), file, 0644)
}
