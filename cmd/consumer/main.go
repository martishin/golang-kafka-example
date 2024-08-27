package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func consumeOrders() {
	kafkaAddr := os.Getenv("KAFKA_ADDR")
	if kafkaAddr == "" {
		kafkaAddr = "localhost:9092"
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "orders"
	}

	consumerGroup := os.Getenv("KAFKA_CONSUMER_GROUP")
	if consumerGroup == "" {
		consumerGroup = "order-consumer-group"
	}

	consumer, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": kafkaAddr,
			"group.id":          consumerGroup,
			"auto.offset.reset": "earliest",
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer func() {
		if closeErr := consumer.Close(); closeErr != nil {
			log.Printf("Error closing consumer: %v", closeErr)
		}
	}()

	err = consumer.Subscribe(kafkaTopic, nil)
	if err != nil {
		log.Printf("Failed to subscribe to topic %s: %v", kafkaTopic, err)
		return
	}

	for {
		msg, msgErr := consumer.ReadMessage(-1) // -1 means the consumer will wait indefinitely for a message
		if msgErr != nil {
			log.Printf("Consumer error: %v (%v)\n", msgErr, msg)
			continue
		}
		// var order models.Order
		// if err := json.Unmarshal(msg.Value, &order); err != nil {
		// 	log.Fatalf("Failed to unmarshal message: %v\n", err)
		// }

		log.Printf("Received order: %+v\n", string(msg.Value))
	}
}

func main() {
	consumeOrders()
}
