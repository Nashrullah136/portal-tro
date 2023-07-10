package export_csv

import "time"

type GetAllRequest struct {
	Username   string     `form:"username"`
	AfterDate  *time.Time `form:"after_date"`
	BeforeDate *time.Time `form:"before_date"`
	Page       int        `form:"page,omitempty"`
	PerPage    int        `form:"perpage"`
}
