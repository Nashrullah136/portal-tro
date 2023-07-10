package audit

import (
	"nashrul-be/crm/repositories"
	"time"
)

type GetAllRequest struct {
	Username string     `form:"username"`
	Object   string     `form:"object"`
	ObjectId string     `form:"object_id"`
	FromDate *time.Time `form:"from" time_format:"2006-01-02"`
	ToDate   *time.Time `form:"to" time_format:"2006-01-02"`
	Page     int        `form:"page,omitempty"`
	PerPage  int        `form:"perpage"`
}

type ExportRequest struct {
	Username string     `form:"username"`
	Object   string     `form:"object"`
	ObjectId string     `form:"object_id"`
	FromDate *time.Time `form:"from" time_format:"2006-01-02"`
	ToDate   *time.Time `form:"to" time_format:"2006-01-02"`
}

type CreateAuditRequest struct {
	Action string `json:"action" binding:"required,printascii"`
}

type PayloadQueue struct {
	RequestID uint
	Query     repositories.AuditQuery
}
