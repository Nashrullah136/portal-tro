package worker

import (
	"encoding/json"
	"github.com/adjust/rmq/v5"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/utils/logutils"
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
		logutils.Get().Printf("Failed to unmarshall payload export csv. error: %s\n", err)
		return
	}
	if err := e.auditRepo.Create(payload); err != nil {
		Reject(&delivery)
		logutils.Get().Printf("Failed to insert new audit to database. error: %s\n", err)
		return
	}
	if err := delivery.Ack(); err != nil {
		logutils.Get().Printf("Failed to write to csv file. error: %s\n", err)
	}
}
