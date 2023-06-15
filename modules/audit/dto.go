package audit

import "time"

type GetAllRequest struct {
	Username   string    `form:"username"`
	Object     string    `form:"object"`
	ObjectId   string    `form:"object_id"`
	AfterDate  time.Time `form:"after_date"`
	BeforeDate time.Time `form:"before_date"`
	Page       int       `form:"page,omitempty"`
	PerPage    int       `form:"perpage"`
}

type CreateAuditRequest struct {
	Action string `json:"action" binding:"required,printascii"`
}
