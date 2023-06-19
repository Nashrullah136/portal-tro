package csv

import (
	"encoding/csv"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

const ExportFolder = "reports"

type FileCsv struct {
	File     *os.File
	Filename string
	Path     string
	Writer   *csv.Writer
}

func NewCSV() (*FileCsv, error) {
	filename := uuid.NewString() + ".csv"
	path := filepath.Join(ExportFolder, filename)
	file, err := os.Create(path)
	writer := csv.NewWriter(file)
	if err != nil {
		return nil, err
	}
	return &FileCsv{
		File:     file,
		Filename: filename,
		Path:     path,
		Writer:   writer,
	}, nil
}

func (c *FileCsv) Write(data []string) error {
	return c.Writer.Write(data)
}

func (c *FileCsv) Finish() {
	c.Writer.Flush()
	c.File.Close()
}
