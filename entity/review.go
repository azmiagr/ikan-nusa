package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ReviewID      int       `json:"review_id" gorm:"type:int;primaryKey;autoIncrement"`
	ReviewContent string    `json:"review_content" gorm:"type:text;not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	ProductID     int       `json:"product_id"`
	UserID        uuid.UUID `json:"user_id"`
}
