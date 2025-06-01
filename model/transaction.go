package model

import "time"

type CheckoutResponse struct {
	TransactionID int                     `json:"transaction_id"`
	TotalPrice    float64                 `json:"total_price"`
	Status        string                  `json:"status"`
	Items         []CheckoutItemsResponse `json:"items"`
	CreatedAt     time.Time
}

type CheckoutItemsResponse struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

type CheckoutRequest struct {
	CartItemsID []int `json:"cart_items_id"`
}
