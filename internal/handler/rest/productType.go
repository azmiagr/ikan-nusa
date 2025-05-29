package rest

import (
	"ikan-nusa/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllTypes(c *gin.Context) {
	res, err := r.service.ProductTypeService.GetAllTypes()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get product types", err)
		return
	}

	response.Success(c, http.StatusOK, "success get product types", res)
}
