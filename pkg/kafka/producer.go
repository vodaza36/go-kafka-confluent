package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Producer class
type Producer struct {
	BootstrapServers string
	Topic            string
	KafkaProducer    *kafka.Producer
}

// NewProducer constructor
func NewProducer(bootstrapServers string, topic string) *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})

	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}

	return &Producer{
		BootstrapServers: bootstrapServers,
		Topic:            topic,
		KafkaProducer:    p,
	}
}

// Send message to Kafka broker
func (p *Producer) Send(value []byte) error {
	err := p.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		return err
	}
	p.KafkaProducer.Flush(15 * 1000)

	return nil
}

// Close the Kafka producer
func (p *Producer) Close() {
	p.KafkaProducer.Close()
}
