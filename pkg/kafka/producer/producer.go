package producer

import (
	"context"
	"log"

	sarama "github.com/IBM/sarama"
)

type producer struct {
	log  *log.Logger
	prod sarama.SyncProducer
}

// NewProducer is producer struct constructor
func NewProducer(logger *log.Logger, brokers []string) (*producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &producer{
		prod: prod,
		log:  logger,
	}, nil
}

// Produce produces message to topic
func (p *producer) Produce(_ context.Context, topic string, data []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}

	partition, offset, err := p.prod.SendMessage(msg)
	if err != nil {
		p.log.Printf("failed to send message in Kafka: %v\n", err.Error())

		return err
	}

	p.log.Printf("message sent to partition %d with offset %d\n", partition, offset)

	return nil
}

func (p *producer) Close() error {
	return p.prod.Close()
}
