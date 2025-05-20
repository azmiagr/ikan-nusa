package model

import "github.com/google/uuid"

type UserParam struct {
	UserID   uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Password string    `json:"-"`
	RoleID   int       `json:"-"`
}

type UserRegister struct {
	Username    string `json:"username" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
}

type UserRegisterResponse struct {
	Token string `json:"token"`
}

type AddAddressAfterRegisterParam struct {
	UserID        uuid.UUID `json:"user_id"`
	DistrictID    int       `json:"district_id" binding:"required"`
	PostalCode    string    `json:"postal_code" binding:"required"`
	AddressDetail string    `json:"address_detail" binding:"required"`
}

type VerifyUser struct {
	UserID  uuid.UUID `json:"user_id" binding:"required"`
	OtpCode string    `json:"otp_code" binding:"required"`
}
