package main

import (
	"log"

	"github.com/vodaza36/go-kafka-confluent/pkg/kafka"
)

func main() {
	c := kafka.NewConsumer("localhost", "myGroup", "myTopic")
	err := c.Listen(myConsumerHandler)

	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
}

func myConsumerHandler(msg []byte) {
	log.Printf("Message: %s", string(msg))
}
