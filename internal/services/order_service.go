package services

import (
	"math/rand"

	"github.com/martishin/golang-kafka-example/internal/models"
)

const (
	customerIDRange = 1000
	costRange       = 10000
	costdenominator = 100
)

type OrderService interface {
	GetOrder() models.Order
}

type RandomOrderService struct{}

func (rog *RandomOrderService) GenerateOrder() models.Order {
	categories := []string{"Electronics", "Books", "Clothing", "Groceries"}
	items := map[string][]string{
		"Electronics": {"Laptop", "Smartphone", "Tablet", "Headphones"},
		"Books":       {"The Go Programming Language", "Kafka: The Definitive Guide", "Clean Code"},
		"Clothing":    {"T-Shirt", "Jeans", "Sweater", "Jacket"},
		"Groceries":   {"Apples", "Bananas", "Carrots", "Potatoes"},
	}

	category := categories[rand.Intn(len(categories))]
	item := items[category][rand.Intn(len(items[category]))]

	return models.Order{
		CustomerID: rand.Intn(customerIDRange), // Random CustomerID between 0-999
		Category:   category,
		Cost:       float64(rand.Intn(costRange)) / costdenominator, // Random cost between 0-100.00
		ItemName:   item,
	}
}
