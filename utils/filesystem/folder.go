package filesystem

import (
	"errors"
	"os"
	"path"
)

var ErrNotExist = errors.New("file not found")

//go:generate mockery --name Folder
type Folder interface {
	Create(filename string) (File, error)
	GetFile(filename string) (File, error)
	GetAllFiles() []File
	GetPath() string
}

type folder struct {
	path string
}

func NewFolder(path string) Folder {
	return folder{path: path}
}

func (f folder) isExist(filename string) bool {
	if _, err := os.Stat(getFilePath(f, filename)); err != nil {
		return false
	}
	return true
}

func (f folder) GetFile(filename string) (File, error) {
	if !f.isExist(filename) {
		return nil, ErrNotExist
	}
	return &file{
		osFile:   nil,
		filename: filename,
		folder:   f,
	}, nil
}

func (f folder) GetAllFiles() (files []File) {
	dirs, _ := os.ReadDir(f.path)
	for _, dir := range dirs {
		if !dir.IsDir() {
			files = append(files, NewFile(dir.Name(), f))
		}
	}
	return
}

func (f folder) Create(filename string) (File, error) {
	osFile, err := os.Create(getFilePath(f, filename))
	if err != nil {
		return nil, err
	}
	return &file{
		osFile:   osFile,
		filename: filename,
		folder:   f,
	}, nil
}

func (f folder) GetPath() string {
	return f.path
}

func getFilePath(folder Folder, filename string) string {
	return path.Join(folder.GetPath(), filename)
}
