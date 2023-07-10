package testutil

import (
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/localtime"
)

func GetSpan() entities.SPAN {
	return entities.SPAN{
		DocumentNumber:      "230081301007314000002",
		DocumentDate:        localtime.Now().Format("2006-01-02"),
		BeneficiaryBankCode: "525119000992",
		StatusCode:          "0002",
		EmailAddress:        "-",
		BeneficiaryAccount:  "1160800947",
		Amount:              "189950000",
		BeneficiaryBank:     "BPD RIAU",
	}
}

func CreateSpan() error {
	db, err := GetConn()
	if err != nil {
		return err
	}
	return db.Model(&entities.SPAN{}).Create(GetSpan()).Error
}

func DeleteSpan() error {
	db, err := GetConn()
	if err != nil {
		return err
	}
	return db.Model(&entities.SPAN{}).Where(map[string]any{
		"DOCUMENTNUMBER": "230081301007314000002",
	}).Delete(&map[string]any{}).Error
}
