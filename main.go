package main

import (
	"context"
	"flag"
)

func main() {

	ctx := context.Background()

	sender := flag.String("s", "", "-s to set the sender")
	receiver := flag.String("r", "", "-r to set the receiver")
	flag.Parse()

	// Kafka consumer and producer types.
	consumerFactory := new(ConsumerFactory)
	consumerType := "kafka"
	consumer, err := consumerFactory.Get(consumerType)
	if err != nil {
		panic(err)
	}

	producerFactory := new(ProducerFactory)
	producerType := "kafka"
	producer, err := producerFactory.Get(producerType)
	if err != nil {
		panic(err)
	}

	chat, err := New(ctx, *sender, *receiver, producer, consumer)
	if err != nil {
		panic(err)
	}

	err = chat.Start()
	if err != nil {
		panic(err)
	}
}
