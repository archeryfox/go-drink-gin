package repository

import (
	"context"

	"gorm.io/gorm"
)

type RecommendationRepository interface {
	Create(ctx context.Context, r *RecommendationModel) error
	ListForUser(ctx context.Context, userID uint, offset, limit int) ([]RecommendationModel, error)
}

type gormRecommendationRepo struct {
	db *gorm.DB
}

func NewGormRecommendationRepository(db *gorm.DB) RecommendationRepository {
	return &gormRecommendationRepo{db: db}
}

func (r *gormRecommendationRepo) Create(ctx context.Context, rec *RecommendationModel) error {
	return r.db.WithContext(ctx).Create(rec).Error
}

func (r *gormRecommendationRepo) ListForUser(ctx context.Context, userID uint, offset, limit int) ([]RecommendationModel, error) {
	var list []RecommendationModel
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
