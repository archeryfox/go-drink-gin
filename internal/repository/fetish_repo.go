package repository

import (
	"context"

	"gorm.io/gorm"
)

type FetishRepository interface {
	Create(ctx context.Context, f *FetishModel) error
	GetByID(ctx context.Context, id uint) (*FetishModel, error)
	List(ctx context.Context, offset, limit int) ([]FetishModel, error)
}

type gormFetishRepo struct {
	db *gorm.DB
}

func NewGormFetishRepository(db *gorm.DB) FetishRepository {
	return &gormFetishRepo{db: db}
}

func (r *gormFetishRepo) Create(ctx context.Context, f *FetishModel) error {
	return r.db.WithContext(ctx).Create(f).Error
}

func (r *gormFetishRepo) GetByID(ctx context.Context, id uint) (*FetishModel, error) {
	var f FetishModel
	if err := r.db.WithContext(ctx).First(&f, id).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *gormFetishRepo) List(ctx context.Context, offset, limit int) ([]FetishModel, error) {
	var list []FetishModel
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
