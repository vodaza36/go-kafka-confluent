package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	err = c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		log.Fatalf("Error subscribing to Kafka topic: %v", err)
	}

	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			log.Fatalf("Error reading message from Kafka topic: %v", err)
		}

		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
}
