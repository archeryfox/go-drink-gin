package integration

import (
	"context"
	"go-gin-hello/internal/repository"
	"go-gin-hello/internal/service"
	"testing"
)

// fake repo for unit tests
type fakeRepo struct {
	store map[uint]*repository.ProductModel
	next  uint
}

func newFakeRepo() *fakeRepo { return &fakeRepo{store: map[uint]*repository.ProductModel{}, next: 1} }

func (f *fakeRepo) Create(ctx context.Context, p *repository.ProductModel) error {
	p.ID = f.next
	f.next++
	f.store[p.ID] = p
	return nil
}
func (f *fakeRepo) GetByID(ctx context.Context, id uint) (*repository.ProductModel, error) {
	v, ok := f.store[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return v, nil
}
func (f *fakeRepo) List(ctx context.Context, offset, limit int) ([]repository.ProductModel, error) {
	var out []repository.ProductModel
	for _, v := range f.store {
		out = append(out, *v)
	}
	return out, nil
}
func (f *fakeRepo) Update(ctx context.Context, p *repository.ProductModel) error {
	f.store[p.ID] = p
	return nil
}
func (f *fakeRepo) Delete(ctx context.Context, id uint) error {
	delete(f.store, id)
	return nil
}

func TestProductService_CreateGet(t *testing.T) {
	repo := newFakeRepo()
	svc := service.NewProductService(repo, nil)
	ctx := context.Background()
	req := service.CreateProductRequest{Name: "Toy", Price: 9.99}
	p, err := svc.Create(ctx, req)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := svc.Get(ctx, p.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != req.Name {
		t.Fatalf("name mismatch")
	}
}
