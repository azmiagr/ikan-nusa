package entity

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	StoreID          uuid.UUID `json:"store_id" gorm:"type:varchar(36);primaryKey"`
	StoreName        string    `json:"store_name" gorm:"type:varchar(90);not null"`
	StoreDescription string    `json:"store_description" gorm:"type:text;not null"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID           uuid.UUID `json:"user_id" `

	Products     []Product     `json:"products" gorm:"foreignKey:StoreID"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:StoreID"`
}
