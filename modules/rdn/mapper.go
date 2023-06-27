package rdn

import "nashrul-be/crm/entities"

func mapUpdateRequestToBriva(request UpdateRequest) entities.Briva {
	return entities.Briva{
		Brivano:  request.Brivano,
		IsActive: request.Active,
	}
}
