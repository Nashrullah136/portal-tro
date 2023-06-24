package csv

import (
	"encoding/csv"
	"github.com/google/uuid"
	"nashrul-be/crm/utils/filesystem"
)

type FileCsv struct {
	File   filesystem.File
	writer *csv.Writer
}

func NewCSV(folder filesystem.Folder) (*FileCsv, error) {
	filename := uuid.NewString() + ".csv"
	file, err := folder.Create(filename)
	if err != nil {
		return nil, err
	}
	osFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(osFile)
	return &FileCsv{
		File:   filesystem.NewFile(filename, folder),
		writer: writer,
	}, nil
}

func (c *FileCsv) Write(data []string) error {
	return c.writer.Write(data)
}

func (c *FileCsv) Finish() {
	c.writer.Flush()
	c.File.Close()
}
