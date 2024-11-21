package env_config

import (
	"fmt"
	"os"
	"strings"

	sarama "github.com/IBM/sarama"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type kafkaConfig struct {
	brokers []string
	groupID string
}

// NewKafkaConfig is kafkaConfig struct constructor
func NewKafkaConfig() (*kafkaConfig, error) {
	brokers := os.Getenv(consts.KafkaBrokersEnvName)
	if len(brokers) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.KafkaBrokersEnvName)
	}

	groupID := os.Getenv(consts.KafkaGroupIDEnvName)
	if len(groupID) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.KafkaGroupIDEnvName)
	}

	return &kafkaConfig{
		brokers: strings.Split(brokers, ","),
		groupID: groupID,
	}, nil
}

// Brokers returns kafka brokers address
func (c *kafkaConfig) Brokers() []string {
	return c.brokers
}

// GroupID returns kafka consumer's group id
func (c *kafkaConfig) GroupID() string {
	return c.groupID
}

// Config returns kafka config
func (c *kafkaConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_6_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}
