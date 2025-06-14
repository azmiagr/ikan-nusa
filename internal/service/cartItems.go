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

type ICartItemsService interface {
	AddToCart(storeID uuid.UUID, cartID int, param *model.AddToCartParam) (*model.AddToCartResponse, error)
	DeleteFromCart(cartItemsID int) error
}

type CartItemsService struct {
	db                  *gorm.DB
	UserRepository      repository.IUserRepository
	CartRepository      repository.ICartRepository
	CartItemsRepository repository.ICartItemsRepository
	ProductRepository   repository.IProductRepository
	StoreRepository     repository.IStoreRepository
}

func NewCartItemsService(userRepository repository.IUserRepository, CartItemsRepository repository.ICartItemsRepository, cartRepository repository.ICartRepository, productRepository repository.IProductRepository, storeRepository repository.IStoreRepository) ICartItemsService {
	return &CartItemsService{
		db:                  mariadb.Connection,
		UserRepository:      userRepository,
		CartItemsRepository: CartItemsRepository,
		CartRepository:      cartRepository,
		ProductRepository:   productRepository,
		StoreRepository:     storeRepository,
	}
}

func (ci *CartItemsService) AddToCart(storeID uuid.UUID, cartID int, param *model.AddToCartParam) (*model.AddToCartResponse, error) {
	var finalCartItems *entity.CartItems

	tx := ci.db.Begin()
	defer tx.Rollback()

	cart, err := ci.CartRepository.GetCartByID(tx, cartID)
	if err != nil {
		return nil, errors.New("users didn'nt have cart")
	}

	store, err := ci.StoreRepository.GetStore(tx, model.StoreParam{
		StoreID: storeID,
	})
	if err != nil {
		return nil, errors.New("store doesn't exist")
	}

	product, err := ci.ProductRepository.GetProduct(model.GetProductParam{
		ProductID: param.ProductID,
	})
	if err != nil {
		return nil, errors.New("product doesn't exist")
	}

	if param.Quantity > product.Stock || param.Quantity <= 0 {
		return nil, errors.New("cannot add more quantity than this")
	}

	existingCartItems, err := ci.CartItemsRepository.GetCartItemsByProductID(tx, product.ProductID)
	if err != nil {
		newCartItems := &entity.CartItems{
			Quantity:  param.Quantity,
			Price:     product.Price,
			CartID:    cart.CartID,
			ProductID: product.ProductID,
		}

		createdItem, err := ci.CartItemsRepository.CreateCartItems(tx, newCartItems)
		if err != nil {
			return nil, err
		}
		finalCartItems = createdItem

	} else {
		existingCartItems.Quantity = param.Quantity
		existingCartItems.Price = product.Price

		updatedItems, err := ci.CartItemsRepository.UpdateCartItems(tx, existingCartItems)
		if err != nil {
			return nil, err
		}
		finalCartItems = updatedItems
	}

	res := &model.AddToCartResponse{
		CartItemsID: finalCartItems.CartItemsID,
		ProductName: product.ProductName,
		StoreName:   store.StoreName,
		Quantity:    finalCartItems.Quantity,
		Price:       product.Price,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (ci *CartItemsService) DeleteFromCart(cartItemsID int) error {
	tx := ci.db.Begin()
	defer tx.Rollback()

	cartItems, err := ci.CartItemsRepository.GetCartItemsByID(tx, cartItemsID)
	if err != nil {
		return errors.New("cart items did'nt exist")
	}

	err = ci.CartItemsRepository.DeleteCartItems(tx, cartItems)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
