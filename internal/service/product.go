package service

import (
	"errors"
	"ikan-nusa/entity"
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
	"ikan-nusa/pkg/database/mariadb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProductService interface {
	AddProduct(userID uuid.UUID, param *model.AddProduct) (*model.AddProductResponse, error)
}

type ProductService struct {
	db                *gorm.DB
	ProductRepository repository.IProductRepository
	StoreRepository   repository.IStoreRepository
}

func NewProductService(productRepository repository.IProductRepository, storeRepository repository.IStoreRepository) IProductService {
	return &ProductService{
		db:                mariadb.Connection,
		ProductRepository: productRepository,
		StoreRepository:   storeRepository,
	}
}

func (p *ProductService) AddProduct(userID uuid.UUID, param *model.AddProduct) (*model.AddProductResponse, error) {
	tx := p.db.Begin()
	defer tx.Rollback()

	store, err := p.StoreRepository.GetStore(tx, model.StoreParam{
		UserID: userID,
	})
	if err != nil {
		return nil, errors.New("user didn't have store")
	}

	product := &entity.Product{
		ProductName:        param.ProductName,
		ProductDescription: param.ProductDescription,
		Price:              param.Price,
		Stock:              param.Stock,
		Category:           param.Category,
		StoreID:            store.StoreID,
	}

	_, err = p.ProductRepository.CreateProduct(tx, product)
	if err != nil {
		return nil, err
	}

	res := &model.AddProductResponse{
		ProductID: product.ProductID,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
