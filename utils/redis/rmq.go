package redisUtils

import (
	"github.com/adjust/rmq/v5"
	"strconv"
	"time"
)

func MakeQueue(conn rmq.Connection, name, tag string, interval int, consumers ...rmq.Consumer) (rmq.Queue, error) {
	queue, err := conn.OpenQueue(name)
	if err != nil {
		return nil, err
	}
	if err := queue.StartConsuming(int64(len(consumers)+1), time.Duration(interval)*time.Second); err != nil {
		return nil, err
	}
	for i, consumer := range consumers {
		if _, err := queue.AddConsumer(tag+strconv.Itoa(i), consumer); err != nil {
			return nil, err
		}
	}
	return queue, nil
}
