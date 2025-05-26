package model

type Ticket struct {
	Destination string  `json:"destination"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
