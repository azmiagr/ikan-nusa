package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/pkg/bcrypt"
	"ikan-nusa/pkg/jwt"
	"ikan-nusa/pkg/supabase"
)

type Service struct {
	UserService      IUserService
	ProductService   IProductService
	CartItemsService ICartItemsService
	ReviewService    IReviewService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) *Service {
	return &Service{
		UserService:      NewUserService(repository.UserRepository, repository.CartRepository, repository.AddressRepository, repository.OtpRepository, repository.StoreRepository, bcrypt, jwtAuth, supabase),
		ProductService:   NewProductService(repository.ProductRepository, repository.StoreRepository),
		CartItemsService: NewCartItemsService(repository.UserRepository, repository.CartItemsRepository, repository.CartRepository, repository.ProductRepository, repository.StoreRepository),
		ReviewService:    NewReviewService(repository.ReviewRepository, repository.UserRepository),
	}
}
