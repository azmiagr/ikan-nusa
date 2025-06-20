package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type IDistrictRepository interface {
	GetAllDistricts() ([]*entity.District, error)
	GetDistrictByCityId(cityID int) ([]*entity.District, error)
}

type DistrictRepository struct {
	db *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) IDistrictRepository {
	return &DistrictRepository{db: db}
}

func (d *DistrictRepository) GetAllDistricts() ([]*entity.District, error) {
	var districts []*entity.District
	err := d.db.Debug().Find(&districts).Error
	if err != nil {
		return nil, err
	}

	return districts, nil
}

func (d *DistrictRepository) GetDistrictByCityId(cityID int) ([]*entity.District, error) {
	var districts []*entity.District
	err := d.db.Debug().Where("city_id = ?", cityID).Find(&districts).Error
	if err != nil {
		return nil, err
	}

	return districts, nil
}
