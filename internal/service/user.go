package service

import (
	"errors"
	"ikan-nusa/entity"
	"ikan-nusa/internal/repository"
	"ikan-nusa/model"
	"ikan-nusa/pkg/bcrypt"
	"ikan-nusa/pkg/database/mariadb"
	"ikan-nusa/pkg/jwt"
	"ikan-nusa/pkg/mail"
	"ikan-nusa/pkg/supabase"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(param *model.UserRegister) (*model.UserRegisterResponse, error)
	AddAddressAfterRegister(param model.AddAddressAfterRegisterParam) error
	VerifyUser(param model.VerifyUser) error
}

type UserService struct {
	db                *gorm.DB
	UserRepository    repository.IUserRepository
	CartRepository    repository.ICartRepository
	AddressRepository repository.IAddressRepository
	OtpRepository     repository.IOtpRepository
	BCrypt            bcrypt.Interface
	JwtAuth           jwt.Interface
	Supabase          supabase.Interface
}

func NewUserService(userRepository repository.IUserRepository, cartRepository repository.ICartRepository, addressRepository repository.IAddressRepository, otpRepository repository.IOtpRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) IUserService {
	return &UserService{
		db:                mariadb.Connection,
		UserRepository:    userRepository,
		CartRepository:    cartRepository,
		AddressRepository: addressRepository,
		OtpRepository:     otpRepository,
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

func (u *UserService) AddAddressAfterRegister(param model.AddAddressAfterRegisterParam) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	user, err := u.UserRepository.GetUser(model.UserParam{
		UserID: param.UserID,
	})
	if err != nil {
		return err
	}

	address := &entity.Address{
		RecipentName:  user.Username,
		PostalCode:    param.PostalCode,
		AddressDetail: param.AddressDetail,
		DistrictID:    param.DistrictID,
		UserID:        user.UserID,
	}

	_, err = u.AddressRepository.CreateAddress(tx, address)
	if err != nil {
		return err
	}

	code := mail.GenerateCode()
	otp := &entity.OtpCode{
		OtpID:  uuid.New(),
		UserID: user.UserID,
		Code:   code,
	}

	err = u.OtpRepository.CreateOtp(tx, otp)
	if err != nil {
		return err
	}

	err = mail.SendEmail(user.Email, "OTP Verification", "Your OTP verification code is "+code+".")
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) VerifyUser(param model.VerifyUser) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	otp, err := u.OtpRepository.GetOtp(tx, model.GetOtp{
		UserID: param.UserID,
	})
	if err != nil {
		return err
	}

	if otp.Code != param.OtpCode {
		return errors.New("invalid otp code")
	}

	expiredTime, err := strconv.Atoi(os.Getenv("EXPIRED_OTP"))
	if err != nil {
		return err
	}

	expiredThreshold := time.Now().UTC().Add(-time.Duration(expiredTime) * time.Minute)
	if otp.UpdatedAt.Before(expiredThreshold) {
		return errors.New("otp expired")
	}

	user, err := u.UserRepository.GetUser(model.UserParam{
		UserID: param.UserID,
	})
	if err != nil {
		return err
	}

	user.StatusAccount = "active"
	err = u.UserRepository.UpdateUser(tx, user)
	if err != nil {
		return err
	}

	err = u.OtpRepository.DeleteOtp(tx, otp)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
