package rest

import (
	"ikan-nusa/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllCities(c *gin.Context) {
	res, err := r.service.CityService.GetAllCities()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get cities data", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get cities data", res)
}
