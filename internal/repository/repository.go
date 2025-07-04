package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository        IUserRepository
	CartRepository        ICartRepository
	AddressRepository     IAddressRepository
	StoreRepository       IStoreRepository
	OtpRepository         IOtpRepository
	ProductRepository     IProductRepository
	CartItemsRepository   ICartItemsRepository
	ReviewRepository      IReviewRepository
	ProductTypeRepository IProductTypeRepository
	TransactionRepository ITransactionRepository
	ProvinceRepository    IProvinceRepository
	CityRepository        ICityRepository
	DistrictRepository    IDistrictRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:        NewUserRepository(db),
		CartRepository:        NewCartRepository(db),
		AddressRepository:     NewAddressRepository(db),
		StoreRepository:       NewStoreRepository(db),
		ProductRepository:     NewProductRepository(db),
		OtpRepository:         NewOtpRepository(db),
		CartItemsRepository:   NewCartItemsRepository(db),
		ReviewRepository:      NewReviewRepository(db),
		ProductTypeRepository: NewProductTypeRepository(db),
		TransactionRepository: NewTransactionRepository(db),
		ProvinceRepository:    NewProvinceRepository(db),
		CityRepository:        NewCityRepository(db),
		DistrictRepository:    NewDistrictRepository(db),
	}
}
