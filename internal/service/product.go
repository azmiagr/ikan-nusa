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
	GetProductsByCategory(category string) ([]*model.GetProductsByCategoryResponse, error)
	GetProductsDetail(productID int) (*model.GetProductsDetailResponse, error)
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

func (p *ProductService) GetProductsByCategory(category string) ([]*model.GetProductsByCategoryResponse, error) {
	tx := p.db.Begin()
	defer tx.Rollback()

	var res []*model.GetProductsByCategoryResponse

	products, err := p.ProductRepository.GetProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	for _, v := range products {
		store, err := p.StoreRepository.GetStore(tx, model.StoreParam{
			StoreID: v.StoreID,
		})
		if err != nil {
			return nil, err
		}

		res = append(res, &model.GetProductsByCategoryResponse{
			ProductName: v.ProductName,
			Price:       v.Price,
			StoreName:   store.StoreName,
		})
	}

	err = tx.Commit().Error

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (p *ProductService) GetProductsDetail(productID int) (*model.GetProductsDetailResponse, error) {
	tx := p.db.Begin()
	defer tx.Rollback()

	product, err := p.ProductRepository.GetProductsDetail(productID)
	if err != nil {
		return nil, err
	}

	store, err := p.StoreRepository.GetStore(tx, model.StoreParam{
		StoreID: product.StoreID,
	})
	if err != nil {
		return nil, err
	}

	res := &model.GetProductsDetailResponse{
		ProductName: product.ProductName,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.ProductDescription,
		Category:    product.Category,
		StoreName:   store.StoreName,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil

}
