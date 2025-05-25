package model

import "github.com/google/uuid"

type AddProduct struct {
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	Category           string  `json:"category"`
	Price              float64 `json:"price"`
	Stock              int     `json:"stock"`
}

type AddProductResponse struct {
	ProductID int `json:"product_id"`
}

type GetProductsByCategoryResponse struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	StoreName   string  `json:"store_name"`
}

type GetProductsDetailResponse struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	StoreName   string  `json:"store_name"`
}

type GetproductsByNameResponse struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	StoreName   string  `json:"store_name"`
}

type GetAllProductsResponse struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	StoreName   string  `json:"store_name"`
}

type GetProductParam struct {
	ProductID int       `json:"-"`
	StoreID   uuid.UUID `json:"-"`
}
