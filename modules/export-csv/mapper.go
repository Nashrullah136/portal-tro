package export_csv

import (
	"nashrul-be/crm/repositories"
)

func mapGetAllRequestToExportCsvQuery(request GetAllRequest) repositories.ExportCsvQuery {
	return repositories.ExportCsvQuery{
		Username:   request.Username,
		AfterDate:  request.AfterDate,
		BeforeDate: request.BeforeDate,
	}
}
