package main

import "github.com/confluentinc/confluent-kafka-go/kafka"

type IProducer interface {
	Produce(m *kafka.Message, deliveryChan chan kafka.Event) error
	Events() chan kafka.Event
}
