package worker

import (
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
)

type Audit struct {
	auditRepo repositories.AuditRepositoryInterface
}

func NewAudit(auditRepo repositories.AuditRepositoryInterface) *ExportCSV {
	return &ExportCSV{
		auditRepo: auditRepo,
	}
}

func (e *Audit) Consume(delivery rmq.Delivery) {
	var payload entities.Audit
	payloadJson := delivery.Payload()
	if err := json.Unmarshal([]byte(payloadJson), &payload); err != nil {
		Reject(&delivery)
		log.Fatalf("Failed to unmarshall payload export csv. error: %s\n", err)
	}
	if err := e.auditRepo.Create(payload); err != nil {
		Reject(&delivery)
		log.Fatalf("Failed to insert new audit to database. error: %s\n", err)
	}
	if err := delivery.Ack(); err != nil {
		log.Printf("Failed to write to csv file. error: %s\n", err)
	}
}
