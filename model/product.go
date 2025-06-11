package model

import "github.com/google/uuid"

type AddProduct struct {
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	Category           string  `json:"category"`
	Price              float64 `json:"price"`
	Stock              int     `json:"stock"`
	ProductTypeID      int     `json:"product_type_id"`
}

type AddProductResponse struct {
	ProductID int `json:"product_id"`
}

type GetProductsByCategoryResponse struct {
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	StoreName     string  `json:"store_name"`
	ImageURL      string  `json:"image_url"`
	ProductTypeID int     `json:"product_type_id"`
}

type GetProductsDetailResponse struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	StoreName   string  `json:"store_name"`
	ProductType string  `json:"product_type"`
}

type GetproductsByNameResponse struct {
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	StoreName     string  `json:"store_name"`
	ImageURL      string  `json:"image_url"`
	ProductTypeID int     `json:"product_type_id"`
}

type GetAllProductsResponse struct {
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	StoreName     string  `json:"store_name"`
	ImageURL      string  `json:"image_url"`
	ProductTypeID int     `json:"product_type_id"`
}

type GetProductsByTypeResponse struct {
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	StoreName     string  `json:"store_name"`
	ImageURL      string  `json:"image_url"`
	ProductTypeID int     `json:"product_type_id"`
}

type GetProductParam struct {
	ProductID int       `json:"-"`
	StoreID   uuid.UUID `json:"-"`
}

type UpdateProductParam struct {
	ProductID int `json:"-"`
	Quantity  int `json:"-"`
}
