package audit

import (
	"nashrul-be/crm/repositories"
	"time"
)

type GetAllRequest struct {
	Username   string    `form:"username"`
	Object     string    `form:"object"`
	ObjectId   string    `form:"object_id"`
	AfterDate  time.Time `form:"after_date" time_format:"01/02/2006"`
	BeforeDate time.Time `form:"before_date" time_format:"01/02/2006"`
	Page       int       `form:"page,omitempty"`
	PerPage    int       `form:"perpage"`
}

type ExportRequest struct {
	Username   string    `form:"username"`
	Object     string    `form:"object"`
	ObjectId   string    `form:"object_id"`
	AfterDate  time.Time `form:"after_date"`
	BeforeDate time.Time `form:"before_date"`
}

type CreateAuditRequest struct {
	Action string `json:"action" binding:"required,printascii"`
}

type PayloadQueue struct {
	RequestID uint
	Query     repositories.AuditQuery
}
