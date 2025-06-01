package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID int       `json:"transaction_id" gorm:"type:int;primaryKey;autoIncrement"`
	TotalPrice    float64   `json:"total_price" gorm:"type:decimal;not null"`
	Status        string    `json:"status" gorm:"type:enum('pending', 'completed');default:'pending'"`
	CreateedAt    time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
	UserID        uuid.UUID `json:"user_id"`
	StoreID       uuid.UUID `json:"store_id"`

	TransactionItems []TransactionItems `json:"transaction_items" gorm:"foreignKey:TransactionID"`
}
