package repository

import (
	"context"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(ctx context.Context, n *NotificationModel) error
	ListForUser(ctx context.Context, userID uint, offset, limit int) ([]NotificationModel, error)
	MarkRead(ctx context.Context, id uint) error
}

type gormNotificationRepo struct {
	db *gorm.DB
}

func NewGormNotificationRepository(db *gorm.DB) NotificationRepository {
	return &gormNotificationRepo{db: db}
}

func (r *gormNotificationRepo) Create(ctx context.Context, n *NotificationModel) error {
	return r.db.WithContext(ctx).Create(n).Error
}

func (r *gormNotificationRepo) ListForUser(ctx context.Context, userID uint, offset, limit int) ([]NotificationModel, error) {
	var list []NotificationModel
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *gormNotificationRepo) MarkRead(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&NotificationModel{}).Where("id = ?", id).Update("read", true).Error
}
