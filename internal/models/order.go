package models

type Order struct {
	CustomerID int     `json:"customer_id"`
	Category   string  `json:"category"`
	Cost       float64 `json:"cost"`
	ItemName   string  `json:"item_name"`
}
