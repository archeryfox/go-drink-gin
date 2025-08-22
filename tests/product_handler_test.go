package tests

import (
	"go-gin-hello/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-gin-hello/internal/repository"
	"go-gin-hello/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupRouterForTest() *gin.Engine {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	err := db.AutoMigrate(&repository.ProductModel{})
	if err != nil {
		return nil
	}
	repo := repository.NewGormProductRepository(db)
	svc := service.NewProductService(repo, nil)
	ph := handlers.NewProductHandler(svc)
	r := gin.Default()
	r.POST("/products", ph.CreateProduct)
	return r
}

func TestCreateProductHandler(t *testing.T) {
	r := setupRouterForTest()
	w := httptest.NewRecorder()
	body := `{"name":"Toy","price":5.5}`
	req, _ := http.NewRequest("POST", "/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d body %s", w.Code, w.Body.String())
	}
}
