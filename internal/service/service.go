package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/pkg/bcrypt"
	"ikan-nusa/pkg/jwt"
	"ikan-nusa/pkg/supabase"
)

type Service struct {
	UserService IUserService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository, repository.CartRepository, repository.AddressRepository, bcrypt, jwtAuth, supabase),
	}
}
