package service

import (
	"ikan-nusa/internal/repository"
	"ikan-nusa/pkg/bcrypt"
	"ikan-nusa/pkg/jwt"
	"ikan-nusa/pkg/supabase"
)

type Service struct {
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) *Service {
	return &Service{}
}
