package audit

import (
	"nashrul-be/crm/repositories"
)

func mapGetAllRequestToAuditQuery(req GetAllRequest) repositories.AuditQuery {
	return repositories.AuditQuery{
		Username: req.Username,
		Object:   req.Object,
		ObjectId: req.ObjectId,
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	}
}

func mapExportRequestToAuditQuery(req ExportRequest) repositories.AuditQuery {
	return repositories.AuditQuery{
		Username: req.Username,
		Object:   req.Object,
		ObjectId: req.ObjectId,
		FromDate: req.AfterDate,
		ToDate:   req.BeforeDate,
	}
}
