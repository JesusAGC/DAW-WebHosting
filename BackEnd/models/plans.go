package models

type Plan struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Details string  `json:"details"`
}
