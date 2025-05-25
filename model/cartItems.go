package model

type AddToCartParam struct {
	Quantity  int `json:"quantity"`
	ProductID int `json:"product_id"`
}

type AddToCartResponse struct {
	CartItemsID int     `json:"cart_items_id"`
	ProductName string  `json:"product_name"`
	StoreName   string  `json:"store_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
