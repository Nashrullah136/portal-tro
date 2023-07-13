package redisUtils

import (
	"github.com/adjust/rmq/v5"
	"github.com/go-co-op/gocron"
	"nashrul-be/crm/utils/logutils"
	"strconv"
	"time"
)

var numberOfQueue = 0

func MakeQueue(conn rmq.Connection, name, tag string, interval int, consumers ...rmq.Consumer) (rmq.Queue, error) {
	queue, err := conn.OpenQueue(name)
	if err != nil {
		return nil, err
	}
	if err := queue.StartConsuming(int64(len(consumers)+1), time.Duration(interval)*time.Second); err != nil {
		return nil, err
	}
	numberOfQueue += 1
	for i, consumer := range consumers {
		if _, err := queue.AddConsumer(tag+strconv.Itoa(i), consumer); err != nil {
			return nil, err
		}
	}
	return queue, nil
}

func CreateCleaner(conn rmq.Connection, scheduler *gocron.Scheduler) error {
	for i := 0; i < numberOfQueue; i++ {
		cleaner := rmq.NewCleaner(conn)
		_, err := scheduler.Every(5).Minute().Do(func() {
			_, err := cleaner.Clean()
			if err != nil {
				logutils.Get().Printf("Failed to clean queue. error: %s\n", err)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}
