package entity

import "github.com/google/uuid"

type User struct {
	UserID        uuid.UUID `json:"user_id" gorm:"type:varchar(36);primaryKey"`
	Username      string    `json:"username" gorm:"type:varchar(70);not null"`
	Email         string    `json:"email" gorm:"type:varchar(70);not null"`
	Password      string    `json:"password" gorm:"type:varchar(70);not null"`
	StatusAccount string    `json:"-" gorm:"type:enum('inactive', 'active');"`
	PhoneNumber   string    `json:"phone_number" gorm:"type:varchar(20);not null"`

	Store        Store         `json:"store" gorm:"foreignKey:UserID"`
	Cart         Cart          `json:"cart" gorm:"foreignKey:UserID"`
	Addresses    []Address     `json:"addresses" gorm:"foreignKey:UserID"`
	Reviews      []Review      `json:"reviews" gorm:"foreignKey:UserID"`
	OtpCode      []OtpCode     `json:"otp_code" gorm:"foreignKey:UserID"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:UserID"`
}
