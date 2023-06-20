package worker

import (
	"github.com/adjust/rmq/v5"
	"log"
)

func Reject(delivery *rmq.Delivery) {
	if err := (*delivery).Reject(); err != nil {
		log.Printf("Fail to reject delivery. error: %s\n", err)
	}
}
