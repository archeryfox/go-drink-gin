package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go-gin-hello/internal/repository"

	"github.com/go-redis/redis/v8"
)

// ProductService defines domain service behaviour
type ProductService interface {
	Create(ctx context.Context, req CreateProductRequest) (*repository.ProductModel, error)
	Get(ctx context.Context, id uint) (*repository.ProductModel, error)
	List(ctx context.Context, offset, limit int) ([]repository.ProductModel, error)
	Update(ctx context.Context, id uint, req UpdateProductRequest) (*repository.ProductModel, error)
	Delete(ctx context.Context, id uint) error
}

// DTOs
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  uint    `json:"category_id"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `json:"category_id"`
}

// Implementation
type productService struct {
	repo  repository.ProductRepository
	cache *redis.Client
}

func NewProductService(r repository.ProductRepository, c *redis.Client) ProductService {
	return &productService{repo: r, cache: c}
}

func (s *productService) cacheKey(id uint) string {
	return fmt.Sprintf("product:%d", id)
}

func (s *productService) Create(ctx context.Context, req CreateProductRequest) (*repository.ProductModel, error) {
	m := &repository.ProductModel{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}
	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	// warm cache
	if s.cache != nil {
		b, _ := json.Marshal(m)
		s.cache.Set(ctx, s.cacheKey(m.ID), b, 10*time.Minute)
	}
	return m, nil
}

func (s *productService) Get(ctx context.Context, id uint) (*repository.ProductModel, error) {
	if s.cache != nil {
		if v, err := s.cache.Get(ctx, s.cacheKey(id)).Result(); err == nil {
			var m repository.ProductModel
			if err := json.Unmarshal([]byte(v), &m); err == nil {
				return &m, nil
			}
		}
	}
	m, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if s.cache != nil {
		b, _ := json.Marshal(m)
		s.cache.Set(ctx, s.cacheKey(m.ID), b, 10*time.Minute)
	}
	return m, nil
}

func (s *productService) List(ctx context.Context, offset, limit int) ([]repository.ProductModel, error) {
	return s.repo.List(ctx, offset, limit)
}

func (s *productService) Update(ctx context.Context, id uint, req UpdateProductRequest) (*repository.ProductModel, error) {
	m, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		m.Name = req.Name
	}
	if req.Description != "" {
		m.Description = req.Description
	}
	if req.Price != 0 {
		m.Price = req.Price
	}
	if req.CategoryID != 0 {
		m.CategoryID = req.CategoryID
	}
	if err := s.repo.Update(ctx, m); err != nil {
		return nil, err
	}
	if s.cache != nil {
		b, _ := json.Marshal(m)
		s.cache.Set(ctx, s.cacheKey(m.ID), b, 10*time.Minute)
	}
	return m, nil
}

func (s *productService) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	if s.cache != nil {
		s.cache.Del(ctx, s.cacheKey(id))
	}
	return nil
}
