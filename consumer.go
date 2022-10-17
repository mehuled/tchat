package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type IConsumer interface {
	Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
}
