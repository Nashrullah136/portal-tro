package worker

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/modules/audit"
	"nashrul-be/crm/repositories"
	csvutils "nashrul-be/crm/utils/csv"
	"nashrul-be/crm/utils/filesystem"
)

type ExportCSV struct {
	auditRepo     repositories.AuditRepositoryInterface
	exportCsvRepo repositories.ExportCsvRepositoryInterface
	folder        filesystem.Folder
}

func NewExportCSV(auditRepo repositories.AuditRepositoryInterface,
	exportCsvRepo repositories.ExportCsvRepositoryInterface,
	folder filesystem.Folder,
) *ExportCSV {
	return &ExportCSV{
		auditRepo:     auditRepo,
		exportCsvRepo: exportCsvRepo,
		folder:        folder,
	}
}

func (e *ExportCSV) Consume(delivery rmq.Delivery) {
	var payload audit.PayloadQueue
	payloadJson := delivery.Payload()
	if err := json.Unmarshal([]byte(payloadJson), &payload); err != nil {
		log.Printf("Failed to unmarshall payload export csv. error: %s\n", err)
		Reject(&delivery)
		return
	}
	exportCsv, err := e.exportCsvRepo.GetById(payload.RequestID)
	if err != nil {
		log.Printf("Failed to get export csv trail data. error: %s\n", err)
		Reject(&delivery)
		return
	}
	exportCsv.Processing()
	if err = e.exportCsvRepo.Update(exportCsv); err != nil {
		log.Printf("Failed to update audit trail data. error: %s\n", err)
		Reject(&delivery)
		return
	}
	audits, err := e.auditRepo.GetAll(context.Background(), payload.Query, 0, 0)
	if err != nil {
		log.Printf("Failed to get audit trail data. error: %s\n", err)
		Reject(&delivery)
		return
	}
	csvFile, err := csvutils.NewCSV(e.folder)
	defer csvFile.Finish()
	if err != nil {
		log.Printf("Failed to create csv file. error: %s\n", err)
		Reject(&delivery)
		return
	}
	if err := csvFile.Write(entities.Audit{}.HeaderCSV()); err != nil {
		log.Printf("Failed to write to csv file. error: %s\n", err)
		Reject(&delivery)
		return
	}
	for _, auditData := range audits {
		if err := csvFile.Write(auditData.CsvRepresentation()); err != nil {
			log.Printf("Failed to write to csv file. error: %s\n", err)
			Reject(&delivery)
			return
		}
	}
	exportCsv.Done(*csvFile)
	if err = e.exportCsvRepo.Update(exportCsv); err != nil {
		log.Printf("Failed to write to csv file. error: %s\n", err)
		Reject(&delivery)
		return
	}
	if err := delivery.Ack(); err != nil {
		log.Printf("Failed to write to csv file. error: %s\n", err)
	}
}
