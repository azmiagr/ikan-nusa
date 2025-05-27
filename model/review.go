package model

type CreateReview struct {
	ReviewContent string `json:"review_content" binding:"required"`
	ProductID     int    `json:"product_id" binding:"required"`
}

type ReviewResponse struct {
	Username      string `json:"username"`
	ReviewContent string `json:"review_content"`
}
