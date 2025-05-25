package rest

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"
	"ikan-nusa/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) AddToCart(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	var param *model.AddToCartParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.CartItemsService.AddToCart(user.Store.StoreID, user.Cart.CartID, param)
	if err != nil {
		if err.Error() == "users didn'nt have cart" {
			response.Error(c, http.StatusInternalServerError, "failed to add products to cart", err)
			return
		} else if err.Error() == "store doesn't exist" {
			response.Error(c, http.StatusBadRequest, "please check the store first", err)
			return
		} else if err.Error() == "product doesn't exist" {
			response.Error(c, http.StatusBadRequest, "please check the product first", err)
			return
		} else if err.Error() == "cannot add more quantity than this" {
			response.Error(c, http.StatusBadRequest, "please check the stock quantity", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to add products to cart", err)
			return
		}

	}

	response.Success(c, http.StatusCreated, "success to add product to cart", res)
}
