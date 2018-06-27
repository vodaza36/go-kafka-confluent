package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Consumer class
type Consumer struct {
	BootstrapServers string
	GroupID          string
	Topic            string
}

// NewConsumer constructor
func NewConsumer(bootstrapServers string, groupID string, topic string) *Consumer {
	return &Consumer{
		BootstrapServers: bootstrapServers,
		GroupID:          groupID,
		Topic:            topic,
	}
}

// Process incoming messages
func (p *Consumer) Process(handler func(msg []byte)) error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": p.BootstrapServers,
		"group.id":          p.GroupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Printf("Error creating Kafka consumer: %v", err)
		return err
	}

	err = c.SubscribeTopics([]string{p.Topic, "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		log.Printf("Error subscribing to Kafka topic: %v", err)
		return err
	}

	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			log.Printf("Error reading message from Kafka topic: %v", err)
			return err
		}

		handler(msg.Value)
	}
}
