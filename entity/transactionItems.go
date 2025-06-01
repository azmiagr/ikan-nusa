package entity

import "time"

type TransactionItems struct {
	TransactionItemsID int       `json:"transaction_items_id" gorm:"type:int;primaryKey;autoIncrement"`
	Quantity           int       `json:"quantity" gorm:"type:int;not null"`
	UnitPrice          float64   `json:"unit_price" gorm:"type:decimal;not null"`
	TotalUnitPrice     float64   `json:"total_unit_price" gorm:"type:decimal;not null"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
	TransactionID      int       `json:"transaction_id"`
	ProductID          int       `json:"product_id"`
}
