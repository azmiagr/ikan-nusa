package entity

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	AddressID     uuid.UUID `json:"address_id" gorm:"type:varchar(36);primaryKey"`
	RecipentName  string    `json:"recipent_name" gorm:"type:varchar(70);not null"`
	PostalCode    string    `json:"postal_code" gorm:"type:varchar(5);not null"`
	AddressDetail string    `json:"address_detail" gorm:"type:varchar(100);not null"`
	UserID        uuid.UUID `json:"user_id"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	District   District `json:"district"`
	DistrictID int      `json:"district_id"`
}
