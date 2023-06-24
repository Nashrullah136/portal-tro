package filesystem

import "os"

type File interface {
	Filename() string
	Path() string
	Open() (*os.File, error)
	Close() error
}

type file struct {
	osFile   *os.File
	filename string
	folder   Folder
}

func NewFile(filename string, folder Folder) File {
	return &file{
		osFile:   nil,
		filename: filename,
		folder:   folder,
	}
}

func (f *file) Filename() string {
	return f.filename
}

func (f *file) Path() string {
	return f.folder.GetPath(f.filename)
}

func (f *file) Open() (*os.File, error) {
	if f.osFile != nil {
		return f.osFile, nil
	}
	return os.Open(f.Path())
}

func (f *file) Close() error {
	if f.osFile == nil {
		return nil
	}
	if err := f.osFile.Close(); err != nil {
		return err
	}
	f.osFile = nil
	return nil
}
