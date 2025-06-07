package model

import "github.com/google/uuid"

type AddressParam struct {
	UserID uuid.UUID `json:"user_id"`
}

type GetCityResponse struct {
	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`
}

type GetDistrictResponse struct {
	DistrictID   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}
