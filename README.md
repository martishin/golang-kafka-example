# Go Kafka Example
## Description
This project demonstrates how to build a simple order processing system using Go and Apache Kafka.   
The system consists of two workers:
* Orders Producer - generates random orders and sends them to a Kafka topic
* Orders Consumer - listens to the Kafka topic and processes the incoming orders

## Running Locally
- Build Docker images: `docker-compose build`
- Start services: `docker-compose up`

## Technologies Used
- [Go](https://go.dev/)
- [Kafka](https://kafka.apache.org/)
- [Confluent Kafka Library](https://github.com/confluentinc/confluent-kafka-go/)
- [Docker](https://www.docker.com/)
