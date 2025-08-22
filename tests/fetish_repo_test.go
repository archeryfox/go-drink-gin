package tests

import (
	"context"
	"go-gin-hello/internal/repository"
	"testing"
)

func TestInMemoryFetishRepo_CreateGet(t *testing.T) {
	repo := repository.NewInMemoryFetishRepo()
	ctx := context.Background()
	f := &repository.FetishModel{Name: "Fursuit"}
	if err := repo.Create(ctx, f); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := repo.GetByID(ctx, f.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != f.Name {
		t.Fatalf("name mismatch")
	}
}
