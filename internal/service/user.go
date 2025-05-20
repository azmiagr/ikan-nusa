package service

import (
	"errors"
	"ikan-nusa/entity"
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
	"ikan-nusa/pkg/bcrypt"
	"ikan-nusa/pkg/database/mariadb"
	"ikan-nusa/pkg/jwt"
	"ikan-nusa/pkg/supabase"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(param *model.UserRegister) (*model.UserRegisterResponse, error)
}

type UserService struct {
	db                *gorm.DB
	UserRepository    repository.IUserRepository
	CartRepository    repository.ICartRepository
	AddressRepository repository.IAddressRepository
	BCrypt            bcrypt.Interface
	JwtAuth           jwt.Interface
	Supabase          supabase.Interface
}

func NewUserService(userRepository repository.IUserRepository, cartRepository repository.ICartRepository, addressRepository repository.IAddressRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) IUserService {
	return &UserService{
		db:                mariadb.Connection,
		UserRepository:    userRepository,
		CartRepository:    cartRepository,
		AddressRepository: addressRepository,
		BCrypt:            bcrypt,
		JwtAuth:           jwtAuth,
		Supabase:          supabase,
	}
}

func (u *UserService) Register(param *model.UserRegister) (*model.UserRegisterResponse, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	var result model.UserRegisterResponse
	_, err := u.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err == nil {
		return nil, errors.New("email already registered")
	}

	hash, err := u.BCrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		UserID:        id,
		Username:      param.Username,
		Email:         param.Email,
		Password:      hash,
		StatusAccount: "inactive",
		PhoneNumber:   param.PhoneNumber,
	}

	_, err = u.UserRepository.CreateUser(tx, user)
	if err != nil {
		return nil, err
	}

	token, err := u.JwtAuth.CreateJWTToken(user.UserID)
	if err != nil {
		return nil, err
	}

	cart := &entity.Cart{
		UserID: user.UserID,
	}

	_, err = u.CartRepository.CreateCart(tx, cart)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	result.Token = token

	return &result, nil

}
