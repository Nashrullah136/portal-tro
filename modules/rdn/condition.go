package rdn

func GetRdnExistCondition() map[string]any {
	return map[string]any{
		"RegStatus": "1003",
		"trxType":   "WIC-EXISTING_CIF",
	}
}

func GetRdnNewCondition() map[string]any {
	return map[string]any{
		"RegStatus": "1003",
		"trxType":   "WIC-NEW_CIF",
	}
}
