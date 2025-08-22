package handlers

import (
	"net/http"
	"strconv"

	repo "go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type ProductFetishHandler struct {
	// will use repository-level functions in future
}

func NewProductFetishHandler() *ProductFetishHandler { return &ProductFetishHandler{} }

// Attach godoc
// @Summary Attach fetish to product
// @Tags product_fetish
// @Accept json
// @Produce json
// @Param product_fetish body repository.ProductFetishModel true "ProductFetish"
// @Success 201 {object} repository.ProductFetishModel
// @Router /product_fetish [post]
func (h *ProductFetishHandler) Attach(c *gin.Context) {
	var pf repo.ProductFetishModel
	if err := c.BindJSON(&pf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pf)
}

func (h *ProductFetishHandler) ListForProduct(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("product_id"))
	c.JSON(http.StatusOK, gin.H{"product_id": pid, "fetishes": []repo.FetishModel{}})
}

// ListForProduct godoc
// @Summary List fetishes for product
// @Tags product_fetish
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {array} repository.FetishModel
// @Router /product_fetish/product/{product_id} [get]
