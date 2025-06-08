package model

import "github.com/google/uuid"

type AddressParam struct {
	UserID uuid.UUID `json:"user_id"`
}

type GetProvinceResponse struct {
	ProvinceID   int    `json:"province_id"`
	ProvinceName string `json:"province_name"`
}

type GetCityResponse struct {
	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`
}

type GetDistrictResponse struct {
	DistrictID   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}
