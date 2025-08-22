package handlers

import (
	"net/http"
	"strconv"

	repo "go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type RecommendationHandler struct {
	repo repo.RecommendationRepository
}

func NewRecommendationHandler(r repo.RecommendationRepository) *RecommendationHandler {
	return &RecommendationHandler{repo: r}
}

// CreateRecommendation godoc
// @Summary Create recommendation
// @Tags recommendations
// @Accept json
// @Produce json
// @Param recommendation body repository.RecommendationModel true "Recommendation"
// @Success 201 {object} repository.RecommendationModel
// @Router /recommendations [post]
func (h *RecommendationHandler) CreateRecommendation(c *gin.Context) {
	var r repo.RecommendationModel
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(c.Request.Context(), &r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

// ListForUser godoc
// @Summary List recommendations for user
// @Tags recommendations
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} repository.RecommendationModel
// @Router /recommendations/user/{user_id} [get]
func (h *RecommendationHandler) ListForUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("user_id"))
	list, err := h.repo.ListForUser(c.Request.Context(), uint(uid), 0, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// swagger:route POST /api/v1/recommendations recommendations createRecommendation
// Create recommendation
//
// Responses:
//   201: RecommendationModel

// swagger:route GET /api/v1/recommendations/user/{user_id} recommendations listRecommendations
// List recommendations for user
//
// Responses:
//   200: []RecommendationModel
