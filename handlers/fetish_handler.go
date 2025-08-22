package handlers

import (
	"net/http"
	"strconv"

	repo "go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type FetishHandler struct {
	repo repo.FetishRepository
}


func NewFetishHandler(r repo.FetishRepository) *FetishHandler {
	return &FetishHandler{repo: r}
}
// CreateFetish godoc
// @Summary Create fetish
// @Tags fetishes
// @Accept json
// @Produce json
// @Param fetish body service.CreateFetishRequest true "Fetish"
// @Success 201 {object} repository.FetishModel
// @Router /fetishes [post]

func (h *FetishHandler) CreateFetish(c *gin.Context) {
	var f repo.FetishModel
	if err := c.BindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(c.Request.Context(), &f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, f)
}

// GetFetish godoc
// @Summary Get fetish
// @Tags fetishes
// @Produce json
// @Param id path int true "Fetish ID"
// @Success 200 {object} repository.FetishModel
// @Failure 404 {object} map[string]string
// @Router /fetishes/{id} [get]
func (h *FetishHandler) GetFetish(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	f, err := h.repo.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, f)
}


// ListFetishes godoc
// @Summary List fetishes
// @Tags fetishes
// @Produce json
// @Success 200 {array} repository.FetishModel
// @Router /fetishes [get]
func (h *FetishHandler) ListFetishes(c *gin.Context) {
	list, err := h.repo.List(c.Request.Context(), 0, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// swagger:route POST /api/v1/fetishes fetishes createFetish
// Create a fetish
//
// Responses:
//   201: FetishModel

// swagger:route GET /api/v1/fetishes fetishes listFetishes
// List fetishes
//
// Responses:
//   200: []FetishModel

// swagger:route GET /api/v1/fetishes/{id} fetishes getFetish
// Get fetish by id
//
// Responses:
//   200: FetishModel
