package tests

import (
	"context"
	"go-gin-hello/internal/repository"
	"testing"
)

func TestInMemoryLikeRepo_CreateCount(t *testing.T) {
	repo := repository.NewInMemoryLikeRepo()
	ctx := context.Background()
	l := &repository.LikeModel{UserID: 1, ProductID: 1}
	if err := repo.Create(ctx, l); err != nil {
		t.Fatalf("create: %v", err)
	}
	cnt, err := repo.CountByProduct(ctx, 1)
	if err != nil {
		t.Fatalf("count: %v", err)
	}
	if cnt != 1 {
		t.Fatalf("expected 1 got %d", cnt)
	}
}
