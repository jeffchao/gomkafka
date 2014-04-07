package gomkafka

import (
	kafka "github.com/Shopify/sarama"
	"log"
	"time"
)

type KafkaConfig struct {
	ClientId string
	Hosts    []string
	Topic    string
}

// Initialize a kafka client and producer based off of KafkaConfig.
func Gomkafka() (*kafka.Client, *kafka.Producer, error) {
	config := KafkaConfig{
		"client_id",
		[]string{"localhost:9092"},
		"monitoring",
	}

	log.Println("initializing Kafka client...")
	client, err := kafka.NewClient(config.ClientId, config.Hosts, &kafka.ClientConfig{MetadataRetries: 1, WaitForElection: 250 * time.Millisecond})
	if err != nil {
		panic(err)
	} else {
		log.Println("kafka client connected")
	}

	log.Println("initializing Kafka producer...")
	producer, err := kafka.NewProducer(client, &kafka.ProducerConfig{RequiredAcks: kafka.WaitForLocal, MaxBufferedBytes: 1, MaxBufferTime: 1})
	if err != nil {
		panic(err)
	}

	return client, producer, nil
}

func Receive() {
}
