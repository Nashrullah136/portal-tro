package testutil

import "nashrul-be/crm/entities"

func GetBriva() entities.Briva {
	return entities.Briva{
		Brivano:  "12121",
		CorpName: "PT. ECHO",
		IsActive: "0",
	}
}

func CreateBriva() error {
	db, err := GetConn()
	if err != nil {
		return err
	}
	return db.Model(&entities.Briva{}).Create(map[string]any{
		"brivano":   "12121",
		"corp_name": "PT. ECHO",
		"IsActive":  "0",
	}).Error
}

func DeleteBriva() error {
	db, err := GetConn()
	if err != nil {
		return err
	}
	return db.Model(&entities.Briva{}).Where(map[string]any{
		"brivano": "12121",
	}).Delete(&map[string]any{}).Error
}
