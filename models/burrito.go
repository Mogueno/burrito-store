package models

type Burrito struct {
	ID    uint    `json:"id"`
	Name  string `json:"name"`
	Size  string `json:"size"`
	Price float64   `json:"price"`
}
