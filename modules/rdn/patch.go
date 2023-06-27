package rdn

import "nashrul-be/crm/entities"

type patch func(rdn entities.RDN) entities.RDN

func PatchRdnExisting(rdn entities.RDN) entities.RDN {
	rdn.IsApproved = "N"
	rdn.IsValidDOBMom = "Y"
	return rdn
}

func PatchRdnNew(rdn entities.RDN) entities.RDN {
	rdn.IsApproved = "N"
	return rdn
}
