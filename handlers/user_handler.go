package handlers

import (
	"net/http"
	"strconv"

	repo "go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo interface {
		// minimal placeholder
	}
}

func NewUserHandler() *UserHandler { return &UserHandler{} }

func (h *UserHandler) CreateUser(c *gin.Context) {
	var u repo.UserModel
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// no DB call yet
	c.JSON(http.StatusCreated, u)
}

// CreateUser godoc
// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body repository.UserModel true "User"
// @Success 201 {object} repository.UserModel
// @Router /users [post]

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetUser godoc
// @Summary Get user by id
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} repository.UserModel
// @Router /users/{id} [get]

// swagger:route POST /api/v1/users users createUser
// Create user
//
// Responses:
//   201: UserModel

// swagger:route GET /api/v1/users/{id} users getUser
// Get user by id
//
// Responses:
//   200: UserModel
