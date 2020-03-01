package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	bootstrapServersVar string = "BOOTSTRAP_SERVERS"
)

var bootstrapServers string = os.Getenv(bootstrapServersVar)

func main() {

	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		panic(fmt.Sprintf("Failed to create producer %s", err))
	}
	defer producer.Close()

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "messages"
	keepRunning := true

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	for keepRunning == true {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			keepRunning = false
		default:
			producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic: &topic, Partition: kafka.PartitionAny},
				Value: []byte("Hello World from Docker"),
			}, nil)
			time.Sleep(1 * time.Second)
		}
	}

}
