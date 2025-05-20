package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type IAddressRepository interface {
}

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) IAddressRepository {
	return &AddressRepository{db: db}
}

func (a *AddressRepository) CreateAddress(tx *gorm.DB, address *entity.Address) (*entity.Address, error) {
	err := tx.Debug().Create(&address).Error
	if err != nil {
		return nil, err
	}

	return address, nil
}
