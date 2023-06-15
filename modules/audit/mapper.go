package audit

import (
	"nashrul-be/crm/repositories"
)

func mapGetAllRequestToAuditQuery(req GetAllRequest) repositories.AuditQuery {
	return repositories.AuditQuery{
		Username:   req.Username,
		Object:     req.Object,
		ObjectId:   req.ObjectId,
		AfterDate:  req.AfterDate,
		BeforeDate: req.BeforeDate,
	}
}
