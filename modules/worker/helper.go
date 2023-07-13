package worker

import (
	"github.com/adjust/rmq/v5"
	"nashrul-be/crm/utils/logutils"
)

func Reject(delivery *rmq.Delivery) {
	if err := (*delivery).Reject(); err != nil {
		logutils.Get().Printf("Fail to reject delivery. error: %s\n", err)
	}
}
