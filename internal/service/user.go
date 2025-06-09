package service

import (
	"errors"
	"fmt"
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
	VerifyUser(param model.VerifyUser) (*model.VerifyUserResponse, error)
	RegisterStore(userID uuid.UUID, param model.RegisterStoreParam) (*model.RegisterStoreResponse, error)
	Login(param model.UserLoginParam) (*model.LoginResponse, error)
	GetUserAddresses(param model.UserParam) ([]*model.GetUserAddresses, error)
	GetUserCartItems(userID uuid.UUID) ([]*model.GetUserCartItemsResponse, error)
	GetNearbyProducts(userID uuid.UUID) ([]*model.GetAllProductsResponse, error)
	Checkout(userID uuid.UUID, param *model.CheckoutRequest) (*model.CheckoutResponse, error)
	GetUser(param model.UserParam) (*entity.User, error)
}

type UserService struct {
	db                    *gorm.DB
	UserRepository        repository.IUserRepository
	CartRepository        repository.ICartRepository
	AddressRepository     repository.IAddressRepository
	StoreRepository       repository.IStoreRepository
	ProductRepository     repository.IProductRepository
	CartItemsRepository   repository.ICartItemsRepository
	TransactionRepository repository.ITransactionRepository
	OtpRepository         repository.IOtpRepository
	BCrypt                bcrypt.Interface
	JwtAuth               jwt.Interface
	Supabase              supabase.Interface
}

func NewUserService(userRepository repository.IUserRepository, cartRepository repository.ICartRepository, addressRepository repository.IAddressRepository, otpRepository repository.IOtpRepository, storeRepository repository.IStoreRepository, productRepository repository.IProductRepository, cartItemsRepository repository.ICartItemsRepository, transactionRepository repository.ITransactionRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) IUserService {
	return &UserService{
		db:                    mariadb.Connection,
		UserRepository:        userRepository,
		CartRepository:        cartRepository,
		AddressRepository:     addressRepository,
		StoreRepository:       storeRepository,
		OtpRepository:         otpRepository,
		ProductRepository:     productRepository,
		CartItemsRepository:   cartItemsRepository,
		TransactionRepository: transactionRepository,
		BCrypt:                bcrypt,
		JwtAuth:               jwtAuth,
		Supabase:              supabase,
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
		AddressID:     uuid.New(),
		RecipentName:  user.Username,
		Label:         param.Label,
		Notes:         param.Notes,
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

func (u *UserService) VerifyUser(param model.VerifyUser) (*model.VerifyUserResponse, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	otp, err := u.OtpRepository.GetOtp(tx, model.GetOtp{
		UserID: param.UserID,
	})
	if err != nil {
		return nil, err
	}

	if otp.Code != param.OtpCode {
		return nil, errors.New("invalid otp code")
	}

	expiredTime, err := strconv.Atoi(os.Getenv("EXPIRED_OTP"))
	if err != nil {
		return nil, err
	}

	expiredThreshold := time.Now().UTC().Add(-time.Duration(expiredTime) * time.Minute)
	if otp.UpdatedAt.Before(expiredThreshold) {
		return nil, errors.New("otp expired")
	}

	user, err := u.UserRepository.GetUser(model.UserParam{
		UserID: param.UserID,
	})
	if err != nil {
		return nil, err
	}

	user.StatusAccount = "active"
	err = u.UserRepository.UpdateUser(tx, user)
	if err != nil {
		return nil, err
	}

	token, err := u.JwtAuth.CreateJWTToken(user.UserID)
	if err != nil {
		return nil, err
	}

	err = u.OtpRepository.DeleteOtp(tx, otp)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	response := &model.VerifyUserResponse{
		Username: user.Username,
		Token:    token,
	}

	return response, nil
}

func (u *UserService) Login(param model.UserLoginParam) (*model.LoginResponse, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	var res model.LoginResponse

	user, err := u.UserRepository.GetUser(model.UserParam{
		Email: param.Email,
	})
	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	err = u.BCrypt.CompareAndHashPassword(user.Password, param.Password)
	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	token, err := u.JwtAuth.CreateJWTToken(user.UserID)
	if err != nil {
		return nil, err
	}

	res.Token = token

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *UserService) RegisterStore(userID uuid.UUID, param model.RegisterStoreParam) (*model.RegisterStoreResponse, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	_, err := u.StoreRepository.GetStore(tx, model.StoreParam{
		StoreName: param.StoreName,
	})
	if err == nil {
		return nil, errors.New("store name already exists")
	}

	store := &entity.Store{
		StoreID:          uuid.New(),
		StoreName:        param.StoreName,
		StoreDescription: param.StoreDescription,
		UserID:           userID,
	}

	_, err = u.StoreRepository.CreateStore(tx, store)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	res := &model.RegisterStoreResponse{
		StoreName:        store.StoreName,
		StoreDescription: store.StoreDescription,
	}

	return res, nil
}

func (u *UserService) GetUserAddresses(param model.UserParam) ([]*model.GetUserAddresses, error) {
	var res []*model.GetUserAddresses

	user, err := u.UserRepository.GetUser(model.UserParam{
		UserID: param.UserID,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range user.Addresses {
		res = append(res, &model.GetUserAddresses{
			DistrictName: v.District.DistrictName,
			CityName:     v.District.City.CityName,
			ProvinceName: v.District.City.Province.ProvinceName,
			Label:        v.Label,
		})
	}

	return res, nil
}

func (u *UserService) GetUserCartItems(userID uuid.UUID) ([]*model.GetUserCartItemsResponse, error) {
	var res []*model.GetUserCartItemsResponse

	tx := u.db.Begin()
	defer tx.Rollback()

	user, err := u.UserRepository.GetUser(model.UserParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	cart, err := u.CartRepository.GetCartByUserID(tx, user.UserID)
	if err != nil {
		return nil, err
	}

	cartItems, err := u.CartItemsRepository.GetCartItemsByCartID(tx, cart.CartID)
	if err != nil {
		return nil, err
	}

	for _, v := range cartItems {
		products, err := u.ProductRepository.GetProduct(model.GetProductParam{
			ProductID: v.ProductID,
		})
		if err != nil {
			return nil, errors.New("product did'nt exist")
		}

		store, err := u.StoreRepository.GetStore(tx, model.StoreParam{
			StoreID: products.StoreID,
		})
		if err != nil {
			return nil, errors.New("store didn'nt exist")
		}

		res = append(res, &model.GetUserCartItemsResponse{
			CartItemsID: v.CartItemsID,
			ProductID:   products.ProductID,
			StoreName:   store.StoreName,
			ImageURL:    products.ImageURL,
			ProductName: products.ProductName,
			Price:       v.Price,
			Quantity:    v.Quantity,
		})
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetNearbyProducts(userID uuid.UUID) ([]*model.GetAllProductsResponse, error) {
	var (
		districtID []int
		res        []*model.GetAllProductsResponse
	)

	tx := u.db.Begin()
	defer tx.Rollback()

	user, err := u.UserRepository.GetUser(model.UserParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range user.Addresses {
		districtID = append(districtID, v.DistrictID)
	}

	products, err := u.ProductRepository.GetProductsByDistricts(districtID)
	if err != nil {
		return nil, err
	}

	for _, v := range products {
		store, err := u.StoreRepository.GetStore(tx, model.StoreParam{
			StoreID: v.StoreID,
		})
		if err != nil {
			return nil, err
		}

		res = append(res, &model.GetAllProductsResponse{
			ProductID:   v.ProductID,
			ProductName: v.ProductName,
			Price:       v.Price,
			StoreName:   store.StoreName,
			ImageURL:    v.ImageURL,
		})
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) Checkout(userID uuid.UUID, param *model.CheckoutRequest) (*model.CheckoutResponse, error) {
	var (
		totalPrice       float64
		transactionItems []*entity.TransactionItems
		items            []model.CheckoutItemsResponse
	)

	tx := u.db.Begin()
	defer tx.Rollback()

	user, err := u.UserRepository.GetUser(model.UserParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	_, err = u.CartRepository.GetCartByUserID(tx, user.UserID)
	if err != nil {
		return nil, err
	}

	cartItems, err := u.CartItemsRepository.GetSelectedCartItems(tx, param.CartItemsID)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("no cart items founds")
	}

	for _, v := range cartItems {
		product, err := u.ProductRepository.GetProduct(model.GetProductParam{
			ProductID: v.ProductID,
		})
		if err != nil {
			return nil, err
		}

		if product.Stock < v.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s. Available: %d, Requested: %d",
				product.ProductName, product.Stock, v.Quantity)
		}

		itemTotal := product.Price * float64(v.Quantity)
		totalPrice += itemTotal

		transactionItems = append(transactionItems, &entity.TransactionItems{
			ProductID:      v.ProductID,
			Quantity:       v.Quantity,
			UnitPrice:      product.Price,
			TotalUnitPrice: itemTotal,
		})
	}

	transaction := &entity.Transaction{
		TotalPrice: totalPrice,
		Status:     "pending",
		UserID:     user.UserID,
	}

	_, err = u.TransactionRepository.CreateTransaction(tx, transaction)
	if err != nil {
		return nil, err
	}

	for _, v := range transactionItems {
		v.TransactionID = transaction.TransactionID
		_, err = u.TransactionRepository.CreateTransactionItems(tx, v)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range cartItems {
		err = u.ProductRepository.UpdateProductStock(tx, &model.UpdateProductParam{
			ProductID: v.ProductID,
			Quantity:  -v.Quantity,
		})
		if err != nil {
			return nil, err
		}
	}

	err = u.CartItemsRepository.DeleteSelectedCartItems(tx, param.CartItemsID)
	if err != nil {
		return nil, err
	}

	for _, v := range transactionItems {
		product, err := u.ProductRepository.GetProduct(model.GetProductParam{
			ProductID: v.ProductID,
		})
		if err != nil {
			return nil, err
		}

		items = append(items, model.CheckoutItemsResponse{
			ProductID:   v.ProductID,
			ProductName: product.ProductName,
			Quantity:    v.Quantity,
			UnitPrice:   v.UnitPrice,
			TotalPrice:  v.TotalUnitPrice,
		})
	}

	res := &model.CheckoutResponse{
		TransactionID: transaction.TransactionID,
		TotalPrice:    totalPrice,
		Status:        transaction.Status,
		Items:         items,
		CreatedAt:     transaction.CreatedAt,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (u *UserService) GetUser(param model.UserParam) (*entity.User, error) {
	return u.UserRepository.GetUser(param)
}
