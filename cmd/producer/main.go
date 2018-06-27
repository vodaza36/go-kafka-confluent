package main

import (
	"github.com/vodaza36/go-kafka-confluent/pkg/kafka"
)

func main() {
	p := kafka.NewProducer("localhost", "myTopic")
	p.Send([]byte("Hello Kafka"))
	defer p.Close()
}
