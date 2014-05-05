package gomkafka

import (
	kafka "github.com/Shopify/sarama"
	"time"
)

/*
KafkaConfig represents the required configuration settings to create a
Kafka client to send requests to a topic in a Kafka cluster.
*/
type KafkaConfig struct {
	ClientID string
	Hosts    []string
	Topic    string
}

/*
Gomkafka will Initialize a kafka client and producer based off of KafkaConfig.
*/
func Gomkafka(config *KafkaConfig) (*kafka.Client, *kafka.Producer, error) {
	client, err := kafka.NewClient(config.ClientID, config.Hosts, &kafka.ClientConfig{MetadataRetries: 1, WaitForElection: 250 * time.Millisecond})
	if err != nil {
		return nil, nil, err
	}

	producer, err := kafka.NewProducer(client, &kafka.ProducerConfig{RequiredAcks: kafka.WaitForLocal, MaxBufferedBytes: 1, MaxBufferTime: 1})
	if err != nil {
		return nil, nil, err
	}

	return client, producer, nil
}
