package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type IProvinceRepository interface {
	GetAllProvince() ([]*entity.Province, error)
}

type ProvinceRepository struct {
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) IProvinceRepository {
	return &ProvinceRepository{db: db}
}

func (p *ProvinceRepository) GetAllProvince() ([]*entity.Province, error) {
	var provinces []*entity.Province
	err := p.db.Debug().Find(&provinces).Error
	if err != nil {
		return nil, err
	}

	return provinces, nil
}
