package span

import "nashrul-be/crm/entities"

func PatchBankRiau(oldData entities.SPAN) entities.SPAN {
	//statuscode = '0000', beneficiarybankcode = '525119000990', emailaddress = beneficiaryaccount, beneficiaryaccount = '00000000000000'
	return entities.SPAN{
		DocumentNumber:      oldData.DocumentNumber,
		BeneficiaryBankCode: "525119000990",
		StatusCode:          "0000",
		EmailAddress:        oldData.BeneficiaryAccount,
		BeneficiaryAccount:  "00000000000000",
	}
}
