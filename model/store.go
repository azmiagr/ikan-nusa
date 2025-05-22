package model

import "github.com/google/uuid"

type StoreParam struct {
	StoreID   uuid.UUID `json:"-"`
	StoreName string    `json:"-"`
}

type RegisterStoreParam struct {
	StoreName        string `json:"store_name" binding:"required"`
	StoreDescription string `json:"store_description" binding:"required"`
}

type RegisterStoreResponse struct {
	StoreName        string `json:"store_name"`
	StoreDescription string `json:"store_description"`
}
