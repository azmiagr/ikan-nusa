package rest

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"
	"ikan-nusa/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(c *gin.Context) {
	var param model.UserRegister
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.UserService.Register(&param)
	if err != nil {
		if err.Error() == "email already registered" {
			response.Error(c, http.StatusConflict, "email already registered", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to register user", err)
			return
		}
	}

	response.Success(c, http.StatusCreated, "success register user", res)

}

func (r *Rest) AddAddressAfterRegister(c *gin.Context) {
	var param model.AddAddressAfterRegisterParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.UserService.AddAddressAfterRegister(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to add address", err)
		return
	}

	response.Success(c, http.StatusCreated, "success add address", nil)
}

func (r *Rest) VerifyUser(c *gin.Context) {
	var param model.VerifyUser
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.UserService.VerifyUser(param)
	if err != nil {
		if err.Error() == "invalid otp code" {
			response.Error(c, http.StatusUnauthorized, "invalid otp code", err)
			return
		} else if err.Error() == "otp expired" {
			response.Error(c, http.StatusUnauthorized, "otp expired", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to verify user", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success verify user", res)
}

func (r *Rest) Login(c *gin.Context) {
	var param model.UserLoginParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.UserService.Login(param)
	if err != nil {
		if err.Error() == "email or password is wrong" {
			response.Error(c, http.StatusBadRequest, "email or password is wrong", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to login user", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success to login user", res)

}

func (r *Rest) RegisterStore(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	var param model.RegisterStoreParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.UserService.RegisterStore(user.UserID, param)
	if err != nil {
		if err.Error() == "store name already exists" {
			response.Error(c, http.StatusConflict, "store name already exists", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to register new store", err)
			return
		}
	}

	response.Success(c, http.StatusCreated, "success register new store", res)

}

func (r *Rest) GetUserAddresses(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	res, err := r.service.UserService.GetUserAddresses(model.UserParam{
		UserID: user.UserID,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get user address", err)
		return
	}

	response.Success(c, http.StatusOK, "success get user addresses", res)
}

func (r *Rest) GetUserCartItems(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	res, err := r.service.UserService.GetUserCartItems(user.UserID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get user cart items", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get user cart items", res)
}

func (r *Rest) GetNearbyProducts(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	res, err := r.service.UserService.GetNearbyProducts(user.UserID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get nearby projects", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get nearby products", res)
}

func (r *Rest) Checkout(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	var param model.CheckoutRequest
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.UserService.Checkout(user.UserID, &param)
	if err != nil {
		if err.Error() == "no cart items founds" {
			response.Error(c, http.StatusBadRequest, "failed to checkout", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to checkout", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success to checkout", res)
}
