package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type ICartItemsRepository interface {
	CreateCartItems(tx *gorm.DB, cartItem *entity.CartItems) (*entity.CartItems, error)
	UpdateCartItems(tx *gorm.DB, cartItems *entity.CartItems) (*entity.CartItems, error)
	GetCartItemsByCartID(tx *gorm.DB, cartID int) ([]*entity.CartItems, error)
	GetCartItemsByProductID(tx *gorm.DB, productID int) (*entity.CartItems, error)
	DeleteCartItems(tx *gorm.DB, cartItems *entity.CartItems) error
}

type CartItemsRepository struct {
	db *gorm.DB
}

func NewCartItemsRepository(db *gorm.DB) ICartItemsRepository {
	return &CartItemsRepository{db: db}
}

func (ci *CartItemsRepository) CreateCartItems(tx *gorm.DB, cartItem *entity.CartItems) (*entity.CartItems, error) {
	err := tx.Debug().Create(&cartItem).Error
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (ci *CartItemsRepository) UpdateCartItems(tx *gorm.DB, cartItems *entity.CartItems) (*entity.CartItems, error) {
	err := tx.Debug().Where("cart_items_id = ?", cartItems.CartItemsID).Updates(&cartItems).Error
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (ci *CartItemsRepository) GetCartItemsByCartID(tx *gorm.DB, cartID int) ([]*entity.CartItems, error) {
	var cartItems []*entity.CartItems
	err := tx.Debug().Where("cart_id = ?", cartID).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (ci *CartItemsRepository) GetCartItemsByProductID(tx *gorm.DB, productID int) (*entity.CartItems, error) {
	var cartItems *entity.CartItems
	err := tx.Debug().Where("product_id = ?", productID).First(&cartItems).Error
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (ci *CartItemsRepository) DeleteCartItems(tx *gorm.DB, cartItems *entity.CartItems) error {
	err := tx.Debug().Delete(&cartItems).Error
	if err != nil {
		return err
	}

	return nil
}
