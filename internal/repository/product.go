package repository

import (
	"context"

	"gorm.io/gorm"
)

// ProductModel represents a product in the shop
// swagger:model ProductModel
type ProductModel struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	CategoryID  uint    `json:"category_id"`
}

// CategoryModel represents a product category
// swagger:model CategoryModel
type CategoryModel struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name"`
}

// ProductRepository defines repository behaviour
type ProductRepository interface {
	Create(ctx context.Context, p *ProductModel) error
	GetByID(ctx context.Context, id uint) (*ProductModel, error)
	List(ctx context.Context, offset, limit int) ([]ProductModel, error)
	Update(ctx context.Context, p *ProductModel) error
	Delete(ctx context.Context, id uint) error
}

// gorm implementation
type gormProductRepo struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ProductRepository {
	return &gormProductRepo{db: db}
}

func (r *gormProductRepo) Create(ctx context.Context, p *ProductModel) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *gormProductRepo) GetByID(ctx context.Context, id uint) (*ProductModel, error) {
	var p ProductModel
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *gormProductRepo) List(ctx context.Context, offset, limit int) ([]ProductModel, error) {
	var list []ProductModel
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *gormProductRepo) Update(ctx context.Context, p *ProductModel) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *gormProductRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&ProductModel{}, id).Error
}
