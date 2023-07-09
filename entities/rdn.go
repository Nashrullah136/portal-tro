package entities

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"nashrul-be/crm/utils/audit"
)

type RDN struct {
	RegistrationID string `gorm:"column:RegistrationID" mapstructure:",omitempty"`
	Rdn            string `gorm:"column:RDN" mapstructure:",omitempty"`
	RegStatus      string `gorm:"column:RegStatus" mapstructure:",omitempty"`
	TrxType        string `gorm:"column:trxType" mapstructure:",omitempty"`
	IsValidDOBMom  string `gorm:"column:isValidDOBMom" mapstructure:",omitempty"`
	IsApproved     string `gorm:"column:isApproved" mapstructure:",omitempty"`
}

func (*RDN) TableName() string {
	return "OpeningAccount"
}

func (r *RDN) LogPresentation() (result map[string]any, err error) {
	if err = mapstructure.Decode(r, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (r *RDN) PrimaryKey() string {
	return fmt.Sprintf("RegistrationID: %s, RDN: %s", r.RegistrationID, r.Rdn)
}

func (r *RDN) EntityName() string {
	return "RDN"
}

func (r *RDN) Copy() audit.Auditor {
	rdn := *r
	return &rdn
}

func (r *RDN) PrimaryFields() map[string]any {
	return map[string]any{
		"RegistrationID": r.RegistrationID,
		"RDN":            r.Rdn,
	}
}
