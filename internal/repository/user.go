package repository

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(tx *gorm.DB, user *entity.User) (*entity.User, error)
	GetUser(param model.UserParam) (*entity.User, error)
	UpdateUser(tx *gorm.DB, user *entity.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	err := tx.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUser(param model.UserParam) (*entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Preload("Addresses.District.City.Province").Preload("Store.Products").Preload("Cart.CartItems").Where(&param).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) UpdateUser(tx *gorm.DB, user *entity.User) error {
	err := tx.Where("user_id = ?", user.UserID).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}
