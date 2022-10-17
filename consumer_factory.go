package main

import "github.com/confluentinc/confluent-kafka-go/kafka"

type ConsumerFactory struct {
}

func (c *ConsumerFactory) Get(consumerType string) (IConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServer,
		"group.id":          "1",
		"auto.offset.reset": "earliest",
		"security.protocol": SecurityProtocol,
		"sasl.mechanisms":   SASLMechanism,
		"sasl.username":     SASLUsername,
		"sasl.password":     SASLPassword,
	})
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
