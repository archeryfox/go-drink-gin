package handlers

import (
	"net/http"
	"strconv"

	repo "go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type LikeHandler struct {
	repo repo.LikeRepository
}

func NewLikeHandler(r repo.LikeRepository) *LikeHandler { return &LikeHandler{repo: r} }

// CreateLike godoc
// @Summary Create like
// @Tags likes
// @Accept json
// @Produce json
// @Param like body repository.LikeModel true "Like"
// @Success 201 {object} repository.LikeModel
// @Router /likes [post]
func (h *LikeHandler) CreateLike(c *gin.Context) {
	var l repo.LikeModel
	if err := c.BindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(c.Request.Context(), &l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, l)
}

// CountByProduct godoc
// @Summary Count likes for product
// @Tags likes
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]int
// @Router /likes/product/{id}/count [get]
func (h *LikeHandler) CountByProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cnt, err := h.repo.CountByProduct(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": cnt})
}

// swagger:route POST /api/v1/likes likes createLike
// Create a like
//
// Responses:
//   201: LikeModel

// swagger:route GET /api/v1/likes/product/{id}/count likes countLikes
// Count likes for product
//
// Responses:
//   200: int
