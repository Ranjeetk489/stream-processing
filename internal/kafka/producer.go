package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ranjeetk489/ecom-service/internal/models"
)

type Producer struct {
	producer *kafka.Producer
	topic    string
}

func NewProducer(bootstrapServers, topic string) (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, err
	}
	return &Producer{producer: p, topic: topic}, nil
}

func (p *Producer) ProduceOrder(order models.Order) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Value:          orderJSON,
	}, nil)
}

func (p *Producer) Close() {
	p.producer.Close()
}
