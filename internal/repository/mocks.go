package repository

import (
	"context"
	"errors"
	"sync"
)

// In-memory product repo for tests
type InMemoryProductRepo struct {
	mu  sync.Mutex
	m   map[uint]ProductModel
	seq uint
}

func NewInMemoryProductRepo() *InMemoryProductRepo {
	return &InMemoryProductRepo{m: make(map[uint]ProductModel)}
}

func (r *InMemoryProductRepo) Create(ctx context.Context, p *ProductModel) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	p.ID = r.seq
	r.m[p.ID] = *p
	return nil
}

func (r *InMemoryProductRepo) GetByID(ctx context.Context, id uint) (*ProductModel, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	v, ok := r.m[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &v, nil
}

func (r *InMemoryProductRepo) List(ctx context.Context, offset, limit int) ([]ProductModel, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var out []ProductModel
	for _, v := range r.m {
		out = append(out, v)
	}
	return out, nil
}

func (r *InMemoryProductRepo) Update(ctx context.Context, p *ProductModel) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.m[p.ID]; !ok {
		return errors.New("not found")
	}
	r.m[p.ID] = *p
	return nil
}

func (r *InMemoryProductRepo) Delete(ctx context.Context, id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.m, id)
	return nil
}

// Minimal in-memory fetish repo
type InMemoryFetishRepo struct {
	mu  sync.Mutex
	m   map[uint]FetishModel
	seq uint
}

func NewInMemoryFetishRepo() *InMemoryFetishRepo {
	return &InMemoryFetishRepo{m: make(map[uint]FetishModel)}
}

func (r *InMemoryFetishRepo) Create(ctx context.Context, f *FetishModel) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	f.ID = r.seq
	r.m[f.ID] = *f
	return nil
}

func (r *InMemoryFetishRepo) GetByID(ctx context.Context, id uint) (*FetishModel, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	v, ok := r.m[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &v, nil
}

func (r *InMemoryFetishRepo) List(ctx context.Context, offset, limit int) ([]FetishModel, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var out []FetishModel
	for _, v := range r.m {
		out = append(out, v)
	}
	return out, nil
}

// Minimal in-memory like repo
type InMemoryLikeRepo struct {
	mu  sync.Mutex
	m   map[uint]LikeModel
	seq uint
}

func NewInMemoryLikeRepo() *InMemoryLikeRepo { return &InMemoryLikeRepo{m: make(map[uint]LikeModel)} }

func (r *InMemoryLikeRepo) Create(ctx context.Context, l *LikeModel) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	l.ID = r.seq
	r.m[l.ID] = *l
	return nil
}

func (r *InMemoryLikeRepo) CountByProduct(ctx context.Context, productID uint) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var cnt int64
	for _, v := range r.m {
		if v.ProductID == productID {
			cnt++
		}
	}
	return cnt, nil
}
