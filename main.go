package main

import (
	"flag"
	"fmt"
	kafka "github.com/Shopify/sarama"
	"github.com/jeffchao/gomkafka/gomkafka"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func run() {
	go handleSignals()
	work()
}

func initConfig() (gomkafka.KafkaConfig, error) {
	config := gomkafka.KafkaConfig{}
	var hostArgs string

	flag.StringVar(&config.ClientId, "c", "", "Kafka client id (REQUIRED)")
	flag.StringVar(&hostArgs, "h", "", "comma-separated list of host addresses with their ports (REQUIRED)")
	flag.StringVar(&config.Topic, "t", "", "Kafka topic (REQUIRED)")

	flag.Parse()

	hosts := strings.Split(hostArgs, ",")
	for _, h := range hosts {
		config.Hosts = append(config.Hosts, h)
	}

	if config.ClientId == "" || len(config.Hosts) == 0 || config.Topic == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	return config, nil
}

func work() {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}

	client, producer, err := gomkafka.Gomkafka(config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	defer producer.Close()

	msg := ""

	for {
		_, err := fmt.Scanf("%s\n", &msg)
		if err != nil {
			return
		}

		err = producer.QueueMessage("monitoring", nil, kafka.StringEncoder(msg))
		if err != nil {
			panic(err)
		}

		select {
		case err = <-producer.Errors():
			if err != nil {
				panic(err)
			}
		default:
			// Perform a noop so sarama can can catch disconnect on the other end.
		}

		time.Sleep(1 * time.Millisecond)
	}
}

func handleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	for s := range signals {
		switch s {
		case syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
			// Catch signals that might terminate the process on behalf all goroutines.
			quit()
		}
	}
}

func quit() {
	// Perform any necessary cleanup here.
	os.Exit(1)
}

func main() {
	run()
}
