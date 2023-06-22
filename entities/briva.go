package entities

import (
	"github.com/mitchellh/mapstructure"
)

type Briva struct {
	Brivano  string `gorm:"primaryKey" mapstructure:",omitempty"`
	CorpName string `mapstructure:",omitempty"`
	IsActive string `gorm:"column:IsActive" mapstructure:",omitempty"`
}

func (b *Briva) Off() {
	b.IsActive = "N"
}

func (b *Briva) On() {
	b.IsActive = "Y"
}

func (b *Briva) LogPresentation() (result map[string]any, err error) {
	if err = mapstructure.Decode(b, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (b *Briva) PrimaryKey() string {
	return b.Brivano
}

func (b *Briva) EntityName() string {
	return "BRIVA"
}

func (b *Briva) Copy() Auditor {
	newBriva := *b
	return &newBriva
}

func (b *Briva) PrimaryFields() map[string]any {
	return map[string]any{
		"brivano": b.Brivano,
	}
}
