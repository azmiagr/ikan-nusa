package repository

import (
	"ikan-nusa/entity"

	"gorm.io/gorm"
)

type IReviewRepository interface {
	CreateReview(tx *gorm.DB, review *entity.Review) (*entity.Review, error)
	GetReviewByProductID(tx *gorm.DB, productID int) ([]*entity.Review, error)
}

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) IReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) CreateReview(tx *gorm.DB, review *entity.Review) (*entity.Review, error) {
	err := tx.Debug().Create(&review).Error
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (r *ReviewRepository) GetReviewByProductID(tx *gorm.DB, productID int) ([]*entity.Review, error) {
	var reviews []*entity.Review
	err := tx.Debug().Where("product_id = ?", productID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
