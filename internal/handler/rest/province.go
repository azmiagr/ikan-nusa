package rest

import (
	"ikan-nusa/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllProvinces(c *gin.Context) {
	res, err := r.service.ProvinceService.GetAllProvinces()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get provinces data", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get province data", res)
}
