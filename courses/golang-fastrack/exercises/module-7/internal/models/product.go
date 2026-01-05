package models

import "time"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	SKU         string    `json:"sku"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}

type SellRequest struct {
	Quantity int `json:"quantity"`
}

type RestockRequest struct {
	Quantity int `json:"quantity"`
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	SKU         string  `json:"sku"`
	Category    string  `json:"category"`
}