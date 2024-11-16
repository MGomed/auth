package auth

import (
	"context"

	msg_bus_errors "github.com/MGomed/auth/internal/storage/message_bus/errors"
	msg_bus_model "github.com/MGomed/auth/internal/storage/message_bus/model"
	kafka "github.com/MGomed/auth/pkg/kafka"
)

type messageBus struct {
	producer kafka.Producer
}

// NewMessageBus is messageBus struct constructor
func NewMessageBus(producer kafka.Producer) *messageBus {
	return &messageBus{
		producer: producer,
	}
}

// SendMessage sending messagee to topic
func (b *messageBus) SendMessage(ctx context.Context, msg *msg_bus_model.Message) error {
	topic, ok := msg_bus_model.MessageTypeToTopic[msg.Type]
	if !ok {
		return msg_bus_errors.ErrUnrecognized
	}

	return b.producer.Produce(ctx, topic, msg.Data)
}
