package entity

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	AddressID     uuid.UUID `json:"address_id" gorm:"type:varchar(36);primaryKey"`
	RecipentName  string    `json:"recipent_name" gorm:"type:varchar(70);not null"`
	Label         string    `json:"label" gorm:"type:varchar(50);not null"`
	AddressDetail string    `json:"address_detail" gorm:"type:varchar(100);not null"`
	Notes         string    `json:"notes" gorm:"type:varchar(30);default:null"`
	UserID        uuid.UUID `json:"user_id"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	District   District `json:"district"`
	DistrictID int      `json:"district_id"`
}
