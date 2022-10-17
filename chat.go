package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
)

type Chat struct {
	ID       string
	Sender   string
	Receiver string
	producer IProducer
	consumer IConsumer
}

func New(ctx context.Context, sender string, receiver string, producer IProducer, consumer IConsumer) (*Chat, error) {
	kafkaAdmin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServer,
		"security.protocol": SecurityProtocol,
		"sasl.mechanisms":   SASLMechanism,
		"sasl.username":     SASLUsername,
		"sasl.password":     SASLPassword,
	})
	if err != nil {
		return nil, err
	}

	metadata, err := kafkaAdmin.GetMetadata(&receiver, false, -1)
	if err != nil {
		return nil, err
	}

	for _, topic := range metadata.Topics {
		if topic.Topic == receiver {
			if topic.Error.Code() == kafka.ErrUnknownTopicOrPart {
				return nil, errors.New("receiver topic does not exist")
			}
		}

		if topic.Topic == sender {
			if topic.Topic == sender {
				if topic.Error.Code() == kafka.ErrUnknownTopicOrPart {
					return nil, errors.New("sender topic does not exist")
				}
			}
		}
	}

	err = consumer.Subscribe(sender, nil)
	if err != nil {
		return nil, err
	}

	return &Chat{
		Sender:   sender,
		Receiver: receiver,
		producer: producer,
		consumer: consumer,
	}, nil
}

func (c *Chat) Start() error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {

		for e := range c.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				{
					if ev.TopicPartition.Error != nil {
						fmt.Printf("failed to produce event to topic. err : %s\n", ev.TopicPartition.Error.Error())
					} else {
						fmt.Println("// message sent\n")
					}
				}
			}
		}
	}()

	//Keep producing messages from STDIN to receiver topic.
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			err := c.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &c.Receiver,
					Partition: kafka.PartitionAny,
				},
				Value: []byte(text),
			}, nil)
			if err != nil {
				return
			}
		}
	}()

	//Keep consuming messages from sender topic and print to STDOUT
	for {
		select {
		case sig := <-sigs:
			{
				fmt.Printf("closing application %s\n", sig.String())
				return nil
			}

		default:
			msg, err := c.consumer.ReadMessage(-1)
			if err != nil {
				//fmt.Println(err.Error())
				continue
			}
			fmt.Println(string(msg.Value))
		}

	}
}

func (c *Chat) End() error {
	return nil
}
