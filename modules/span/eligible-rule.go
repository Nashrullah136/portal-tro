package span

import (
	"nashrul-be/crm/entities"
)

func eligibleForPatchBankRiau(span entities.SPAN) bool {
	return span.EmailAddress == "-"
}
