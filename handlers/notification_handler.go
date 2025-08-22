package handlers

import (
	"net/http"
	"strconv"

	repo "go-gin-hello/internal/repository"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	repo repo.NotificationRepository
}

func NewNotificationHandler(r repo.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{repo: r}
}

// CreateNotification godoc
// @Summary Create notification
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body repository.NotificationModel true "Notification"
// @Success 201 {object} repository.NotificationModel
// @Router /notifications [post]
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var n repo.NotificationModel
	if err := c.BindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.Create(c.Request.Context(), &n); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, n)
}

// ListForUser godoc
// @Summary List notifications for user
// @Tags notifications
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} repository.NotificationModel
// @Router /notifications/user/{user_id} [get]
func (h *NotificationHandler) ListForUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("user_id"))
	list, err := h.repo.ListForUser(c.Request.Context(), uint(uid), 0, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// MarkRead godoc
// @Summary Mark notification as read
// @Tags notifications
// @Param id path int true "Notification ID"
// @Success 204
// @Router /notifications/{id}/read [put]
func (h *NotificationHandler) MarkRead(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.MarkRead(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// swagger:route POST /api/v1/notifications notifications createNotification
// Create notification
//
// Responses:
//   201: NotificationModel

// swagger:route GET /api/v1/notifications/user/{user_id} notifications listNotifications
// List notifications for user
//
// Responses:
//   200: []NotificationModel

// swagger:route PUT /api/v1/notifications/{id}/read notifications markNotificationRead
// Mark notification as read
//
// Responses:
//   204:
