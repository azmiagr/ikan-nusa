package rest

import (
	"ikan-nusa/entity"
	"ikan-nusa/model"
	"ikan-nusa/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) AddReview(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	var param model.CreateReview
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	res, err := r.service.ReviewService.AddReview(user.UserID, param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to add review", err)
		return
	}

	response.Success(c, http.StatusOK, "success add review", res)
}
