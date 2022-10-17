package main

import (
	"testing"
)

func TestProducerFactory_Get(t *testing.T) {
	type args struct {
		producerType string
	}
	tests := []struct {
		name    string
		args    args
		want    IProducer
		wantErr bool
	}{
		{
			name:    "test producer get",
			args:    args{producerType: ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProducerFactory{}
			got, err := p.Get(tt.args.producerType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
