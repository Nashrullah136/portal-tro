package entities

import (
	"nashrul-be/crm/utils/csv"
	"time"
)

type ExportCsv struct {
	ID        uint
	User      string
	Status    string
	Filename  string
	CreatedAt time.Time
}

func InitExportCsv(username string) ExportCsv {
	return ExportCsv{
		User:   username,
		Status: "INITIATING",
	}
}

func (ec *ExportCsv) Processing() {
	ec.Status = "PROCESSING"
}

func (ec *ExportCsv) Done(csv csv.FileCsv) {
	ec.Status = "SUCCESS"
	ec.Filename = csv.Filename
}
