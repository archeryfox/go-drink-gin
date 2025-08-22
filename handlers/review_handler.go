package handlers

import (
	"net/http"
	"strconv"

	"go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct{}

func NewReviewHandler() *ReviewHandler {
	return &ReviewHandler{}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var r repository.ReviewModel
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// no-op in current implementation; in production this would call repository
	c.JSON(http.StatusCreated, r)
}

// CreateReview godoc
// @Summary Create review
// @Tags reviews
// @Accept json
// @Produce json
// @Param review body repository.ReviewModel true "Review"
// @Success 201 {object} repository.ReviewModel
// @Router /reviews [post]

func (h *ReviewHandler) ListReviews(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("product_id"))
	c.JSON(http.StatusOK, gin.H{"product_id": pid, "reviews": []repository.ReviewModel{}})
}

// ListReviews godoc
// @Summary List reviews for product
// @Tags reviews
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {array} repository.ReviewModel
// @Router /reviews/product/{product_id} [get]

// swagger:route POST /api/v1/reviews reviews createReview
// Create review
//
// Responses:
//   201: ReviewModel

// swagger:route GET /api/v1/reviews/product/{product_id} reviews listReviews
// List reviews for product
//
// Responses:
//   200: []ReviewModel
