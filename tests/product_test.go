package tests

import (
	"context"
	"go-gin-hello/internal/repository"
	"testing"
)

func TestInMemoryProductRepo(t *testing.T) {
	repo := repository.NewInMemoryProductRepo()
	ctx := context.Background()
	p := &repository.ProductModel{Name: "Test", Price: 1.23}
	if err := repo.Create(ctx, p); err != nil {
		t.Fatalf("create: %v", err)
	}
	if p.ID == 0 {
		t.Fatalf("expected id set")
	}
	got, err := repo.GetByID(ctx, p.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != p.Name {
		t.Fatalf("expected name %s got %s", p.Name, got.Name)
	}
	if err := repo.Delete(ctx, p.ID); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, err := repo.GetByID(ctx, p.ID); err == nil {
		t.Fatalf("expected not found")
	}
}
