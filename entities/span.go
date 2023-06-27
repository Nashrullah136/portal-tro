package entities

import "github.com/mitchellh/mapstructure"

type SPAN struct {
	DocumentNumber      string `gorm:"primaryKey,column:DOCUMENTNUMBER" mapstructure:",omitempty"`
	DocumentDate        string `gorm:"column:DOCUMENTDATE" mapstructure:",omitempty"`
	BeneficiaryBankCode string `gorm:"column:BENEFICIARYBANKCODE" mapstructure:",omitempty"`
	StatusCode          string `gorm:"column:STATUSCODE" mapstructure:",omitempty"`
	EmailAddress        string `gorm:"column:EMAILADDRESS" mapstructure:",omitempty"`
	BeneficiaryAccount  string `gorm:"column:BENEFICIARYACCOUNT" mapstructure:",omitempty"`
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

func (r *SPAN) Copy() Auditor {
	SPAN := *r
	return &SPAN
}

func (r *SPAN) PrimaryFields() map[string]any {
	return map[string]any{
		"documentnumber": r.DocumentNumber,
	}
}
