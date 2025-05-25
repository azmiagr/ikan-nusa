package repository

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(tx *gorm.DB, product *entity.Product) (*entity.Product, error)
	UpdateProduct(tx *gorm.DB, product *entity.Product) (*entity.Product, error)
	GetProductsByCategory(category string) ([]*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
	GetProductsByStoreID(storeID uuid.UUID) ([]*entity.Product, error)
	GetProductsDetail(productID int) (*entity.Product, error)
	GetProduct(param model.GetProductParam) (*entity.Product, error)
	GetProductsByName(productName string) ([]*entity.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) CreateProduct(tx *gorm.DB, product *entity.Product) (*entity.Product, error) {
	err := tx.Debug().Create(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepository) UpdateProduct(tx *gorm.DB, product *entity.Product) (*entity.Product, error) {
	err := tx.Debug().Where("product_id = ?", product.ProductID).Updates(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepository) GetAllProducts() ([]*entity.Product, error) {
	var products []*entity.Product
	err := p.db.Debug().Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetProductsByCategory(category string) ([]*entity.Product, error) {
	var products []*entity.Product
	err := p.db.Debug().Where("category = ?", category).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetProductsByStoreID(storeID uuid.UUID) ([]*entity.Product, error) {
	var products []*entity.Product
	err := p.db.Debug().Where("store_id = ?", storeID).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetProductsDetail(productID int) (*entity.Product, error) {
	var product entity.Product
	err := p.db.Debug().Where("product_id = ?", productID).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepository) GetProductsByName(productName string) ([]*entity.Product, error) {
	var products []*entity.Product
	err := p.db.Debug().Where("product_name LIKE ?", "%"+productName+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetProduct(param model.GetProductParam) (*entity.Product, error) {
	var product *entity.Product
	err := p.db.Debug().Where(&param).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}
