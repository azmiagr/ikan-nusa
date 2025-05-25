package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository      IUserRepository
	CartRepository      ICartRepository
	AddressRepository   IAddressRepository
	StoreRepository     IStoreRepository
	OtpRepository       IOtpRepository
	ProductRepository   IProductRepository
	CartItemsRepository ICartItemsRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:      NewUserRepository(db),
		CartRepository:      NewCartRepository(db),
		AddressRepository:   NewAddressRepository(db),
		StoreRepository:     NewStoreRepository(db),
		ProductRepository:   NewProductRepository(db),
		OtpRepository:       NewOtpRepository(db),
		CartItemsRepository: NewCartItemsRepository(db),
	}
}
