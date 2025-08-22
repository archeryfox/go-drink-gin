package handlers

import (
	"net/http"
	"strconv"

	repo "go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	// reuse no repo for now
}

func NewCategoryHandler() *CategoryHandler { return &CategoryHandler{} }

// CreateCategory godoc
// @Summary Create category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body repository.CategoryModel true "Category"
// @Success 201 {object} repository.CategoryModel
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var cat repo.CategoryModel
	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cat)
}

// GetCategory godoc
// @Summary Get category by id
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} repository.CategoryModel
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// swagger:route POST /api/v1/categories categories createCategory
// Create category
//
// Responses:
//   201: CategoryModel

// swagger:route GET /api/v1/categories/{id} categories getCategory
// Get category by id
//
// Responses:
//   200: CategoryModel
