package model

import "github.com/google/uuid"

type AddressParam struct {
	UserID uuid.UUID `json:"user_id"`
}
