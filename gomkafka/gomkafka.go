package gomkafka

import (
	kafka "github.com/Shopify/sarama"
	"time"
)

type KafkaConfig struct {
	ClientId string
	Hosts    []string
	Topic    string
}

// Initialize a kafka client and producer based off of KafkaConfig.
func Gomkafka(config KafkaConfig) (*kafka.Client, *kafka.Producer, error) {
	client, err := kafka.NewClient(config.ClientId, config.Hosts, &kafka.ClientConfig{MetadataRetries: 1, WaitForElection: 250 * time.Millisecond})
	if err != nil {
		panic(err)
	}

	producer, err := kafka.NewProducer(client, &kafka.ProducerConfig{RequiredAcks: kafka.WaitForLocal, MaxBufferedBytes: 1, MaxBufferTime: 1})
	if err != nil {
		panic(err)
	}

	return client, producer, nil
}

func Receive() {
}
