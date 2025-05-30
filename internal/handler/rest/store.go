package rest

import (
	"ikan-nusa/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetStoreDetail(c *gin.Context) {
	storeName := c.Param("store_name")

	res, err := r.service.StoreService.GetStoreDetail(storeName)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get store details", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get store details", res)
}
