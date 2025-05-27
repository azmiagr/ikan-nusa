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

type IReviewService interface {
	AddReview(userID uuid.UUID, param model.CreateReview) (*model.ReviewResponse, error)
}

type ReviewService struct {
	db                *gorm.DB
	ReviewRepository  repository.IReviewRepository
	UserRepository    repository.IUserRepository
	ProductRepository repository.IProductRepository
}

func NewReviewService(reviewRepository repository.IReviewRepository, userRepository repository.IUserRepository) IReviewService {
	return &ReviewService{
		db:               mariadb.Connection,
		ReviewRepository: reviewRepository,
		UserRepository:   userRepository,
	}
}

func (r *ReviewService) AddReview(userID uuid.UUID, param model.CreateReview) (*model.ReviewResponse, error) {
	tx := r.db.Begin()
	defer tx.Rollback()

	user, err := r.UserRepository.GetUser(model.UserParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	product, err := r.ProductRepository.GetProduct(model.GetProductParam{
		ProductID: param.ProductID,
	})
	if err != nil {
		return nil, errors.New("this product doesn't exist")
	}

	review := &entity.Review{
		ReviewContent: param.ReviewContent,
		ProductID:     product.ProductID,
		UserID:        user.UserID,
	}

	_, err = r.ReviewRepository.CreateReview(tx, review)
	if err != nil {
		return nil, err
	}

	res := &model.ReviewResponse{
		Username:      user.Username,
		ReviewContent: review.ReviewContent,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil

}
