package models

type Order struct {
	ID        uint         `json:"id"`
	Items     []OrderItem `json:"items"`
	TotalCost float64        `json:"total_cost"`
}

type OrderItem struct {
	ID 	 uint     `json:"id"`
	Burrito  Burrito `json:"burrito"`
	Quantity uint     `json:"quantity"`
}
