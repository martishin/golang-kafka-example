package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/martishin/golang-kafka-example/internal/models"
	"github.com/martishin/golang-kafka-example/internal/services"
)

const (
	flushTimeout  = 15000
	orderChanSize = 100
)

func generateOrders(orderChan chan<- models.Order) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	orderGenerator := &services.RandomOrderService{}

	for range ticker.C {
		orderChan <- orderGenerator.GenerateOrder()
	}
	close(orderChan)
}

func produceOrders(orderChan <-chan models.Order) {
	kafkaAddr := os.Getenv("KAFKA_ADDR")
	if kafkaAddr == "" {
		kafkaAddr = "localhost:9092"
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "orders"
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaAddr})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	for order := range orderChan {
		orderJSON, jsonErr := json.Marshal(order)
		if jsonErr != nil {
			log.Printf("Failed to marshal order: %v\n", jsonErr)
			continue
		}

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
			Value:          orderJSON,
		}

		produceErr := producer.Produce(message, nil)
		if produceErr != nil {
			log.Printf("Failed to produce message: %v", produceErr)
			continue
		}

		log.Println("Produced order:", string(orderJSON))
	}
	producer.Flush(flushTimeout)
}

func main() {
	orderChan := make(chan models.Order, orderChanSize)

	go generateOrders(orderChan)
	go produceOrders(orderChan)

	select {}
}
