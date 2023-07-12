package span

import "nashrul-be/crm/entities"

func mapUpdateRequestToSpan(request UpdateRequest) entities.SPAN {
	return entities.SPAN{
		DocumentNumber: request.DocumentNumber,
	}
}

func mapSpanToPresentation(span entities.SPAN) Presentation {
	return Presentation{
		DocumentNumber:      span.DocumentNumber,
		DocumentDate:        span.DocumentDate,
		BeneficiaryBankCode: span.BeneficiaryBankCode,
		StatusCode:          span.StatusCode,
		EmailAddress:        span.EmailAddress,
		BeneficiaryAccount:  span.BeneficiaryAccount,
		Amount:              span.Amount,
		BeneficiaryBank:     span.BeneficiaryBank,
		IsPatched:           eligibleForPatchBankRiau(span),
	}
}
