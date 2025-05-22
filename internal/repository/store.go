package repository

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"

	"gorm.io/gorm"
)

type IStoreRepository interface {
	CreateStore(tx *gorm.DB, store *entity.Store) (*entity.Store, error)
	UpdateStore(tx *gorm.DB, store *entity.Store) (*entity.Store, error)
	GetStore(tx *gorm.DB, param model.StoreParam) (*entity.Store, error)
}

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) IStoreRepository {
	return &StoreRepository{db: db}
}

func (s *StoreRepository) CreateStore(tx *gorm.DB, store *entity.Store) (*entity.Store, error) {
	err := s.db.Debug().Create(&store).Error
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreRepository) UpdateStore(tx *gorm.DB, store *entity.Store) (*entity.Store, error) {
	err := s.db.Debug().Where("store_id = ?", store.StoreID).Updates(&store).Error
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreRepository) GetStore(tx *gorm.DB, param model.StoreParam) (*entity.Store, error) {
	store := &entity.Store{}
	err := s.db.Debug().Where(&param).First(&store).Error
	if err != nil {
		return nil, err
	}

	return store, nil
}
