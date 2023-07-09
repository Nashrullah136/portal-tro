package audit

import "time"

type Result struct {
	DateTime   time.Time
	Username   string
	Action     string
	Entity     string
	EntityID   string
	DataBefore string
	DataAfter  string
}
