package rest

import (
	"ikan-nusa/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllDistricts(c *gin.Context) {
	res, err := r.service.DistrictService.GetAllDistricts()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get districts data", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get districts data", res)
}

func (r *Rest) GetDistrictByCityId(c *gin.Context) {
	idStr := c.Param("city_id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid city id", err)
		return
	}

	res, err := r.service.DistrictService.GetDistrictByCityId(idInt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get districts", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get districts data", res)
}
