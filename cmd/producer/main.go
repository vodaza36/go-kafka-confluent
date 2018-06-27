package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})

	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}

	defer p.Close()

	topic := "myTopic"

	for _, value := range []string{"Hello", "Kafka"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(value),
		}, nil)
	}

	p.Flush(15 * 1000)
}
