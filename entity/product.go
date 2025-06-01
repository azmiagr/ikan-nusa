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
	Category           string    `json:"category" gorm:"type:enum('air tawar', 'air laut', 'lain-lain');"`
	ImageURL           string    `json:"image_url" gorm:"text"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	StoreID            uuid.UUID `json:"store_id"`
	ProductTypeID      int       `json:"product_type_id"`

	Reviews          []Review           `json:"reviews" gorm:"foreignKey:ProductID"`
	CartItems        []CartItems        `json:"cart_items" gorm:"foreignKey:ProductID"`
	TransactionItems []TransactionItems `json:"transaction_items" gorm:"foreignKey:ProductID"`
}
