package rest

import (
	"ikan-nusa/pkg/response"
	"net/http"
	"strconv"

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

func (r *Rest) GetCitiesByProvinceID(c *gin.Context) {
	idStr := c.Param("province_id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid province id", err)
		return
	}

	res, err := r.service.CityService.GetCitiesByProvinceID(idInt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get cities data", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get cities data", res)
}
