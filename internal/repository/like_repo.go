package repository

import (
	"context"

	"gorm.io/gorm"
)

type LikeRepository interface {
	Create(ctx context.Context, l *LikeModel) error
	CountByProduct(ctx context.Context, productID uint) (int64, error)
}

type gormLikeRepo struct {
	db *gorm.DB
}

func NewGormLikeRepository(db *gorm.DB) LikeRepository {
	return &gormLikeRepo{db: db}
}

func (r *gormLikeRepo) Create(ctx context.Context, l *LikeModel) error {
	return r.db.WithContext(ctx).Create(l).Error
}

func (r *gormLikeRepo) CountByProduct(ctx context.Context, productID uint) (int64, error) {
	var cnt int64
	if err := r.db.WithContext(ctx).Model(&LikeModel{}).Where("product_id = ?", productID).Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}
