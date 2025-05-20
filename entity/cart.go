package entity

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	CartID    int       `json:"cart_id" gorm:"type:int;primaryKey;autoIncrement"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
