package entity

import "time"

type CartItems struct {
	CartItemsID int       `json:"cart_items_id" gorm:"type:int;primaryKey;autoIncrement"`
	Quantity    int       `json:"quantity" gorm:"type:int;not null"`
	Price       float64   `json:"price" gorm:"type:decimal;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
}
