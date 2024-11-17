package kafka

import (
	"context"

	consumer "github.com/MGomed/auth/pkg/kafka/consumer"
)

//go:generate mockgen -destination=./mocks/kafka_mock.go -package=mocks -source=interfaces.go

// Consumer is kafka consumer wrapper interface
type Consumer interface {
	Consume(ctx context.Context, topic string, handler consumer.Handler) error
	Close() error
}

// Producer is kafka producer wrapper interface
type Producer interface {
	Produce(ctx context.Context, topic string, msg []byte) error
	Close() error
}
