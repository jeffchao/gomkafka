package gomkafka

import (
	"testing"
  "github.com/Shopify/sarama"
)

func TestGomkafka(t *testing.T) {
  mb := sarama.NewMockBroker(t, 1)
  mb.Returns(new(sarama.MetadataResponse))

	config := &KafkaConfig{
    "client_id",
		[]string{mb.Addr()},
		"foo_topic",
	}
	_, _, err := Gomkafka(config)
	if err != nil {
		t.Errorf("%+v", err)
	}


  config = &KafkaConfig{
    "client_id",
    []string{"localhost:8080"},
    "foo_topic",
  }
  _, _, err = Gomkafka(config)
  if err == nil {
    t.Errorf("expected an error, got <nil>")
  }
}
