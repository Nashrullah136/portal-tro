package span

import "nashrul-be/crm/entities"

func mapUpdateRequestToSpan(request UpdateRequest) entities.SPAN {
	return entities.SPAN{
		DocumentNumber: request.DocumentNumber,
	}
}
