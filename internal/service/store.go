package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
	"ikan-nusa/pkg/database/mariadb"

	"gorm.io/gorm"
)

type IStoreService interface {
	GetStoreDetail(storeName string) (*model.GetStoreDetailResponse, error)
}

type StoreService struct {
	db                *gorm.DB
	StoreRepoository  repository.IStoreRepository
	UserRepository    repository.IUserRepository
	AddressRepository repository.IAddressRepository
}

func NewStoreService(storeRepository repository.IStoreRepository, userRepository repository.IUserRepository, addressRepository repository.IAddressRepository) IStoreService {
	return &StoreService{
		db:                mariadb.Connection,
		StoreRepoository:  storeRepository,
		UserRepository:    userRepository,
		AddressRepository: addressRepository,
	}
}

func (s *StoreService) GetStoreDetail(storeName string) (*model.GetStoreDetailResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	store, err := s.StoreRepoository.GetStore(tx, model.StoreParam{
		StoreName: storeName,
	})
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.GetUser(model.UserParam{
		UserID: store.UserID,
	})
	if err != nil {
		return nil, err
	}

	address, err := s.AddressRepository.GetAddress(tx, model.AddressParam{
		UserID: user.UserID,
	})
	if err != nil {
		return nil, err
	}

	res := &model.GetStoreDetailResponse{
		StoreName: store.StoreName,
		CityName:  address.District.City.CityName,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
