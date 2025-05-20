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
