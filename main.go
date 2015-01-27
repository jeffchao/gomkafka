package main

import (
	"bufio"
	"fmt"
	kafka "github.com/Shopify/sarama"
	"github.com/jeffchao/gomkafka/gomkafka"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func run() error {
	go handleSignals()
	err := work()
	if err != nil {
		return err
	}

	return nil
}

func initConfig() (*gomkafka.KafkaConfig, error) {
	config := &gomkafka.KafkaConfig{}

	if len(os.Args) != 4 {
		printUsage()
		os.Exit(2)
	}

	config.ClientID = os.Args[1]
	for _, h := range strings.Split(os.Args[2], ",") {
		config.Hosts = append(config.Hosts, h)
	}
	config.Topic = os.Args[3]

	if config.ClientID == "" || len(config.Hosts) == 0 || config.Topic == "" {
		printUsage()
		os.Exit(2)
	}

	return config, nil
}

func printUsage() {
	fmt.Printf("Usage: gomkafka id hosts topic\n")
	fmt.Printf("\tid\tKafka client id (REQUIRED)\n")
	fmt.Printf("\thosts\tComma-separated list of host:port pairs (REQUIRED)\n")
	fmt.Printf("\ttopic\tKafka topic (REQUIRED)\n")
}

func work() error {
	config, err := initConfig()
	if err != nil {
		return err
	}

	client, producer, err := gomkafka.Gomkafka(config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	defer producer.Close()

	in := bufio.NewReader(os.Stdin)

	for {
		msg, err := in.ReadString('\n')
		if err != nil {
			return err
		}

		err = producer.SendMessage("monitoring", nil, kafka.StringEncoder(msg))
		if err != nil {
			return err
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
	err := run()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
