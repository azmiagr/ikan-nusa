package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type ICityRepository interface {
	GetAllCities() ([]*entity.City, error)
	GetCitiesByProvinceID(provinceID int) ([]*entity.City, error)
}

type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) ICityRepository {
	return &CityRepository{db: db}
}

func (c *CityRepository) GetAllCities() ([]*entity.City, error) {
	var cities []*entity.City
	err := c.db.Debug().Find(&cities).Error
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func (c *CityRepository) GetCitiesByProvinceID(provinceID int) ([]*entity.City, error) {
	var cities []*entity.City
	err := c.db.Debug().Where("province_id = ?", provinceID).Find(&cities).Error
	if err != nil {
		return nil, err
	}

	return cities, nil
}
