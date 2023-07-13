package rdn

import "nashrul-be/crm/entities"

func mapUpdateRequestToRDN(request UpdateRequest) entities.Briva {
	return entities.Briva{
		Brivano:  request.Brivano,
		IsActive: request.Active,
	}
}
