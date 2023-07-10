package repositories

import "time"

type AuditQuery struct {
	Username string
	Object   string
	ObjectId string
	FromDate *time.Time
	ToDate   *time.Time
}

type ExportCsvQuery struct {
	Username   string
	AfterDate  *time.Time
	BeforeDate *time.Time
}
