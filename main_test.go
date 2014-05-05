package main

import (
	"github.com/Shopify/sarama"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	os.Args = nil
	os.Args = append(os.Args, "gomkafka")
	os.Args = append(os.Args, "client_id")
	os.Args = append(os.Args, "localhost:8080")
	os.Args = append(os.Args, "foo_topic")
	_, err := initConfig()
	if err != nil {
		t.Errorf("error: failed to initialize gomkafka config.")
	}
}

func TestWork(t *testing.T) {
	mb := sarama.NewMockBroker(t, 1)
	mb.Returns(new(sarama.MetadataResponse))

	os.Args = nil
	os.Args = append(os.Args, "gomkafka")
	os.Args = append(os.Args, "client_id")
	os.Args = append(os.Args, mb.Addr())
	os.Args = append(os.Args, "foo_topic")

	err := work()
	// Test empty input case.
	if err == nil {
		t.Errorf("error. %+v", err)
	}
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("error. got: %d, expected: %d.", actual, expected)
	}
}
