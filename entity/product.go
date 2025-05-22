package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID          int       `json:"product_id" gorm:"type:int;primaryKey;autoIncrement"`
	ProductName        string    `json:"product_name" gorm:"type:varchar(65);not null"`
	ProductDescription string    `json:"product_description" gorm:"type:text;not null"`
	Price              float64   `json:"price" gorm:"type:decimal;not null"`
	Stock              int       `json:"stock" gorm:"type:int;not null"`
	Category           string    `json:"category" gorm:"enum('air tawar', 'air laut', 'air payau');"`
	ImageURL           string    `json:"image_url" gorm:"text"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	StoreID            uuid.UUID `json:"store_id"`
}
