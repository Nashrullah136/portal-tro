package filesystem

import (
	"errors"
	"os"
	"path"
)

type Folder interface {
	IsExist(filename string) bool
	Create(filename string) (File, error)
	Remove(filename string) error
	GetPath(filename string) string
}

type folder struct {
	path string
}

func NewFolder(path string) Folder {
	return folder{path: path}
}

func (f folder) IsExist(filename string) bool {
	if _, err := os.Stat(f.GetPath(filename)); err != nil {
		return false
	}
	return true
}

func (f folder) Create(filename string) (File, error) {
	osFile, err := os.Create(f.GetPath(filename))
	if err != nil {
		return nil, err
	}
	return &file{
		osFile:   osFile,
		filename: filename,
		folder:   f,
	}, nil
}

func (f folder) Remove(filename string) error {
	if !f.IsExist(filename) {
		return errors.New("filename doesn't exist")
	}
	return os.Remove(f.GetPath(filename))
}

func (f folder) GetPath(filename string) string {
	return path.Join(f.path, filename)
}
