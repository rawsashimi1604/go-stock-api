package models

type Stock struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Company string  `json:"company"`
}
