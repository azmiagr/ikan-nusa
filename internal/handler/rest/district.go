package rest

import (
	"ikan-nusa/pkg/response"
	"net/http"

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
