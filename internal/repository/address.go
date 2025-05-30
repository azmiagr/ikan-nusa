package repository

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"

	"gorm.io/gorm"
)

type IAddressRepository interface {
	CreateAddress(tx *gorm.DB, address *entity.Address) (*entity.Address, error)
	GetAddress(tx *gorm.DB, param model.AddressParam) (*entity.Address, error)
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

func (a *AddressRepository) GetAddress(tx *gorm.DB, param model.AddressParam) (*entity.Address, error) {
	var address *entity.Address
	err := tx.Debug().Preload("District.City.Province").Where(&param).First(&address).Error
	if err != nil {
		return nil, err
	}

	return address, nil
}
