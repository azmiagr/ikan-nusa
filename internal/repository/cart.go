package repository

import (
	"ikan-nusa/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(tx *gorm.DB, cart *entity.Cart) (*entity.Cart, error)
	GetCartByID(tx *gorm.DB, cartID int) (*entity.Cart, error)
	GetCartByUserID(tx *gorm.DB, userID uuid.UUID) (*entity.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{db: db}
}

func (c *CartRepository) CreateCart(tx *gorm.DB, cart *entity.Cart) (*entity.Cart, error) {
	err := tx.Debug().Create(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (c *CartRepository) GetCartByID(tx *gorm.DB, cartID int) (*entity.Cart, error) {
	var cart *entity.Cart
	err := tx.Debug().Where("cart_id = ?", cartID).Find(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (c *CartRepository) GetCartByUserID(tx *gorm.DB, userID uuid.UUID) (*entity.Cart, error) {
	var cart *entity.Cart
	err := tx.Debug().Where("user_id = ?", userID).Find(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}
