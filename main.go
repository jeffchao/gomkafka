package main

import (
	"fmt"
	"github.com/jeffchao/gomkafka/gomkafka"
	"log"
  "time"
	"os"
	"os/signal"
	"syscall"
)

func run() {
  go handleSignals()
  work()
}

func work() {
	log.Println("waiting for rsyslog input...")
	rsyslog := ""

	for {
		_, err := fmt.Scanf("%s\n", &rsyslog)

		if err != nil {
			return
		}

		fmt.Println(rsyslog)
    time.Sleep(1 * time.Millisecond)
	}
}

func handleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	for s := range signals {
		switch s {
		case syscall.SIGINT, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
			quit()
		}
	}
}

func quit() {
	log.Println("Waiting for cleanup...")
  // Perform any necessary cleanup here.
	log.Println("Exiting")
	os.Exit(1)
}

func main() {
	client, producer, err := gomkafka.Gomkafka()
	if err != nil {
		panic(err)
	}
	defer client.Close()
	defer producer.Close()

  run()
}
