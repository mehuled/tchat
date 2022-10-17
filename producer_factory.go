package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProducerFactory struct {
}

func (p *ProducerFactory) Get(producerType string) (IProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServer,
		"security.protocol": SecurityProtocol,
		"sasl.mechanisms":   SASLMechanism,
		"sasl.username":     SASLUsername,
		"sasl.password":     SASLPassword,
	})
	if err != nil {
		return nil, err
	}

	producer.Events()

	return producer, nil
}
