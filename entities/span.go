package entities

import (
	"github.com/mitchellh/mapstructure"
	"nashrul-be/crm/utils/auditUtils"
)

type SPAN struct {
	DocumentNumber      string `gorm:"column:DOCUMENTNUMBER;primaryKey" mapstructure:",omitempty"`
	DocumentDate        string `gorm:"column:DOCUMENTDATE" mapstructure:",omitempty"`
	BeneficiaryBankCode string `gorm:"column:BENEFICIARYBANKCODE" mapstructure:",omitempty"`
	StatusCode          string `gorm:"column:STATUSCODE" mapstructure:",omitempty"`
	EmailAddress        string `gorm:"column:EMAILADDRESS" mapstructure:",omitempty"`
	BeneficiaryAccount  string `gorm:"column:BENEFICIARYACCOUNT" mapstructure:",omitempty"`
	Amount              string `gorm:"column:AMOUNT"  mapstructure:",omitempty"`
	BeneficiaryBank     string `gorm:"column:BENEFICIARYBANK" mapstructure:",omitempty"`
}

func (*SPAN) TableName() string {
	return "SPANTRANSACTION"
}

func (r *SPAN) LogPresentation() (result map[string]any, err error) {
	if err = mapstructure.Decode(r, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (r *SPAN) PrimaryKey() string {
	return r.DocumentNumber
}

func (r *SPAN) EntityName() string {
	return "SPAN"
}

func (r *SPAN) Copy() auditUtils.Auditor {
	SPAN := *r
	return &SPAN
}

func (r *SPAN) PrimaryFields() map[string]any {
	return map[string]any{
		"documentnumber": r.DocumentNumber,
	}
}
