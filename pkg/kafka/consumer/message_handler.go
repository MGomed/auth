package consumer

import (
	"context"
	"log"

	sarama "github.com/IBM/sarama"
)

// Handler defines message handling func
type Handler func(ctx context.Context, msg *sarama.ConsumerMessage) error

// GroupHandler is consumer group handler
type GroupHandler struct {
	log        *log.Logger
	msgHandler Handler
}

// NewGroupHandler is GroupHandler struct constructor
func NewGroupHandler(logger *log.Logger) *GroupHandler {
	return &GroupHandler{
		log: logger,
	}
}

// Setup is started at the beginning of a new session before ConsumeClaim is called
func (c *GroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of the session life after all ConsumeClaim goroutines are completed
func (c *GroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim should start a consumer message loop ConsumerGroupClaim().
// After the Messages() channel is closed, the handler must complete the processing
func (c *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				c.log.Printf("message channel was closed\n")
				return nil
			}

			c.log.Printf("message claimed: value = %s, timestamp = %v, topic = %s\n",
				string(message.Value), message.Timestamp, message.Topic)

			err := c.msgHandler(session.Context(), message)
			if err != nil {
				c.log.Printf("error handling message: %v\n", err)
				continue
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			c.log.Printf("session context done\n")
			return nil
		}
	}
}
