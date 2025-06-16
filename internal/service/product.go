package service

import (
	"errors"
	"ikan-nusa/entity"
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
	"ikan-nusa/pkg/database/mariadb"
	"ikan-nusa/pkg/supabase"
	"mime/multipart"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProductService interface {
	AddProduct(userID uuid.UUID, param *model.AddProduct) (*model.AddProductResponse, error)
	GetProductsByCategory(category string) ([]*model.GetProductsByCategoryResponse, error)
	GetProductsDetail(productID int) (*model.GetProductsDetailResponse, error)
	GetProductsByName(productName string) ([]*model.GetproductsByNameResponse, error)
	GetAllProducts() ([]*model.GetAllProductsResponse, error)
	GetProductsByType(typeID int) ([]*model.GetProductsByTypeResponse, error)
	UploadPhoto(productID int, file *multipart.FileHeader) (string, error)
}

type ProductService struct {
	db                    *gorm.DB
	ProductTypeRepository repository.IProductTypeRepository
	ProductRepository     repository.IProductRepository
	StoreRepository       repository.IStoreRepository
	Supabase              supabase.Interface
}

func NewProductService(productTypeRepository repository.IProductTypeRepository, productRepository repository.IProductRepository, storeRepository repository.IStoreRepository, supabase supabase.Interface) IProductService {
	return &ProductService{
		db:                    mariadb.Connection,
		ProductTypeRepository: productTypeRepository,
		ProductRepository:     productRepository,
		StoreRepository:       storeRepository,
		Supabase:              supabase,
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
		ProductTypeID:      param.ProductTypeID,
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
			ProductID:     v.ProductID,
			ProductName:   v.ProductName,
			Price:         v.Price,
			ImageURL:      v.ImageURL,
			Stock:         v.Stock,
			StoreName:     store.StoreName,
			ProductTypeID: v.ProductTypeID,
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

	productType, err := p.ProductTypeRepository.GetTypeByID(product.ProductTypeID)
	if err != nil {
		return nil, err
	}

	res := &model.GetProductsDetailResponse{
		ProductID:   product.ProductID,
		ProductName: product.ProductName,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.ProductDescription,
		Category:    product.Category,
		ImageURL:    product.ImageURL,
		StoreName:   store.StoreName,
		ProductType: productType.Type,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (p *ProductService) GetProductsByName(productName string) ([]*model.GetproductsByNameResponse, error) {
	tx := p.db.Begin()
	defer tx.Rollback()

	var res []*model.GetproductsByNameResponse

	products, err := p.ProductRepository.GetProductsByName(productName)
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

		res = append(res, &model.GetproductsByNameResponse{
			ProductID:     v.ProductID,
			ProductName:   v.ProductName,
			Price:         v.Price,
			Stock:         v.Stock,
			ImageURL:      v.ImageURL,
			StoreName:     store.StoreName,
			ProductTypeID: v.ProductTypeID,
		})
	}

	return res, nil
}

func (p *ProductService) GetAllProducts() ([]*model.GetAllProductsResponse, error) {
	tx := p.db.Begin()
	defer tx.Rollback()

	var res []*model.GetAllProductsResponse

	products, err := p.ProductRepository.GetAllProducts()
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

		res = append(res, &model.GetAllProductsResponse{
			ProductID:     v.ProductID,
			ProductName:   v.ProductName,
			Price:         v.Price,
			Stock:         v.Stock,
			StoreName:     store.StoreName,
			ImageURL:      v.ImageURL,
			ProductTypeID: v.ProductTypeID,
		})
	}

	return res, err
}

func (p *ProductService) UploadPhoto(productID int, file *multipart.FileHeader) (string, error) {
	tx := p.db.Begin()
	defer tx.Rollback()

	product, err := p.ProductRepository.GetProduct(model.GetProductParam{
		ProductID: productID,
	})
	if err != nil {
		return "", errors.New("product not found")
	}

	photoURL, err := p.Supabase.UploadFile(file)
	if err != nil {
		return "", err
	}

	product.ImageURL = photoURL

	_, err = p.ProductRepository.UpdateProduct(tx, product)
	if err != nil {
		return "", err
	}

	err = tx.Commit().Error
	if err != nil {
		return "", err
	}

	return photoURL, nil
}

func (p *ProductService) GetProductsByType(typeID int) ([]*model.GetProductsByTypeResponse, error) {
	tx := p.db.Begin()
	defer tx.Rollback()

	var res []*model.GetProductsByTypeResponse

	product, err := p.ProductRepository.GetProductsByType(typeID)
	if err != nil {
		return nil, err
	}

	for _, v := range product {
		store, err := p.StoreRepository.GetStore(tx, model.StoreParam{
			StoreID: v.StoreID,
		})
		if err != nil {
			return nil, err
		}

		res = append(res, &model.GetProductsByTypeResponse{
			ProductID:     v.ProductID,
			ProductName:   v.ProductName,
			Price:         v.Price,
			Stock:         v.Stock,
			StoreName:     store.StoreName,
			ImageURL:      v.ImageURL,
			ProductTypeID: v.ProductTypeID,
		})
	}

	return res, nil
}
