package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ranjeetk489/ecom-service/internal/generator"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "ecom-service-producer",
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	topic := "orders"

	for {
		order := generator.GenerateFakeOrder()
		orderJSON, err := json.Marshal(order)
		if err != nil {
			log.Printf("Failed to marshal order: %v", err)
			continue
		}

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          orderJSON,
		}, nil)

		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		}

		time.Sleep(time.Second) // Produce a message every second
	}
}
