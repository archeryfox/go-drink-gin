package tests

import (
	"go-gin-hello/internal/repository"
	"testing"
)

func TestCategoryModel_Simple(t *testing.T) {
	// basic struct sanity check
	c := repository.CategoryModel{Name: "Toys"}
	if c.Name == "" {
		t.Fatalf("expected name set")
	}
}
