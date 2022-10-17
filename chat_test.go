package main

import (
	"golang.org/x/net/context"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	cFactory := &ConsumerFactory{}
	consumer, err := cFactory.Get("")
	if err != nil {
		panic(err)
	}

	pFactory := &ProducerFactory{}
	producer, err := pFactory.Get("")

	type args struct {
		ctx      context.Context
		sender   string
		receiver string
		producer IProducer
		consumer IConsumer
	}
	tests := []struct {
		name    string
		args    args
		want    *Chat
		wantErr bool
	}{
		{
			name: "check metadata",
			args: args{
				ctx:      nil,
				sender:   "sender_1",
				receiver: "receiver_1",
				producer: producer,
				consumer: consumer,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ctx, tt.args.sender, tt.args.receiver, tt.args.producer, tt.args.consumer)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
