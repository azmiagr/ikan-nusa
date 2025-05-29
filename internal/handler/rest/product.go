package rest

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"
	"ikan-nusa/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) AddProduct(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	var param model.AddProduct
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.ProductService.AddProduct(user.UserID, &param)
	if err != nil {
		if err.Error() == "user didn't have store" {
			response.Error(c, http.StatusBadRequest, "please register your store first", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to add product", err)
			return
		}
	}

	response.Success(c, http.StatusCreated, "success add new product", res)
}

func (r *Rest) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")

	res, err := r.service.ProductService.GetProductsByCategory(category)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get products", err)
		return
	}

	response.Success(c, http.StatusOK, "success get products", res)
}

func (r *Rest) GetProductsDetail(c *gin.Context) {
	productID := c.Param("product_id")

	idInt, err := strconv.Atoi(productID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format", err)
		return
	}

	res, err := r.service.ProductService.GetProductsDetail(idInt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get products detail", err)
		return
	}

	response.Success(c, http.StatusOK, "success get products detail", res)
}

func (r *Rest) GetProductsByName(c *gin.Context) {
	productName := c.Query("product")

	res, err := r.service.ProductService.GetProductsByName(productName)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get products", err)
		return
	}

	response.Success(c, http.StatusOK, "success get products", res)
}

func (r *Rest) GetAllProducts(c *gin.Context) {
	res, err := r.service.ProductService.GetAllProducts()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get all products", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get all products", res)
}

func (r *Rest) UploadPhoto(c *gin.Context) {
	photo, err := c.FormFile("photo")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to upload photo", err)
		return
	}

	productID := c.Param("product_id")
	idInt, err := strconv.Atoi(productID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid product id", err)
		return
	}

	res, err := r.service.ProductService.UploadPhoto(idInt, photo)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to update photo", err)
		return
	}

	response.Success(c, http.StatusOK, "success to update photo", res)
}

func (r *Rest) GetProductsByType(c *gin.Context) {
	typeID := c.Param("product_type_id")
	idInt, err := strconv.Atoi(typeID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid product type id", err)
		return
	}

	res, err := r.service.ProductService.GetProductsByType(idInt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get products", err)
		return
	}

	response.Success(c, http.StatusOK, "succcess to get products", res)
}
