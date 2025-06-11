package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type IProductTypeRepository interface {
	GetAllTypes() ([]*entity.ProductType, error)
	GetTypeByID(typeID int) (*entity.ProductType, error)
}

type ProductTypeRepository struct {
	db *gorm.DB
}

func NewProductTypeRepository(db *gorm.DB) IProductTypeRepository {
	return &ProductTypeRepository{db: db}
}

func (p *ProductTypeRepository) GetAllTypes() ([]*entity.ProductType, error) {
	var productTypes []*entity.ProductType
	err := p.db.Debug().Find(&productTypes).Error
	if err != nil {
		return nil, err
	}

	return productTypes, nil
}

func (p *ProductTypeRepository) GetTypeByID(typeID int) (*entity.ProductType, error) {
	var productType *entity.ProductType
	err := p.db.Debug().Where("product_type_id=?", typeID).First(&productType).Error
	if err != nil {
		return nil, err
	}

	return productType, nil
}
