package entities

import "time"

type Audit struct {
	ID         int        `json:"id,omitempty"`
	DateTime   *time.Time `json:"date_time"`
	Username   string     `json:"username,omitempty"`
	Action     string     `json:"action,omitempty"`
	Entity     string     `json:"entity,omitempty"`
	EntityID   string     `json:"entity_id,omitempty"`
	DataBefore string     `json:"data_before,omitempty"`
	DataAfter  string     `json:"data_after,omitempty"`
}

func (a Audit) HeaderCSV() []string {
	return []string{
		"username",
		"date and time",
		"activity",
		"object",
		"object id",
		"data before",
		"data after",
	}
}

func (a Audit) CsvRepresentation() []string {
	return []string{
		a.Username,
		a.DateTime.String(),
		a.Action,
		a.Entity,
		a.EntityID,
		a.DataBefore,
		a.DataAfter,
	}
}
