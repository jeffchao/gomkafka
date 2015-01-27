package gomkafka

import (
	"github.com/Shopify/sarama"
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
func Gomkafka(config *KafkaConfig) (*sarama.Client, *sarama.SimpleProducer, error) {
	client, err := sarama.NewClient(config.ClientID, config.Hosts, sarama.NewClientConfig())
	if err != nil {
		return nil, nil, err
	}

	producer, err := sarama.NewSimpleProducer(client, nil)
	if err != nil {
		return nil, nil, err
	}

	return client, producer, nil
}
