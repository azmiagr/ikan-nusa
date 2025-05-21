package rest

import (
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

	err = r.service.UserService.VerifyUser(param)
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

	response.Success(c, http.StatusOK, "success verify user", nil)
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
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to login user", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success to login user", res)

}
