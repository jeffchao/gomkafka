package main

import (
	"fmt"
	kafka "github.com/Shopify/sarama"
	"github.com/jeffchao/gomkafka/gomkafka"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func run() {
	go handleSignals()
	work()
}

func work() {
	client, producer, err := gomkafka.Gomkafka()
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

		log.Println("sending")
		err = producer.SendMessage("monitoring", nil, kafka.StringEncoder(msg))
		log.Println("sent")
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Millisecond)
	}
}

func handleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	for s := range signals {
		switch s {
		case syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
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
