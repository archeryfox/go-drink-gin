package integration

import (
	"context"
	"testing"
	"time"

	repo "go-gin-hello/internal/repository"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestSeedDB(t *testing.T) {
	// open in-memory sqlite (shared cache so multiple connections see same data)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}

	// migrate all models
	if err := db.AutoMigrate(&repo.ProductModel{}, &repo.CategoryModel{}, &repo.UserModel{}, &repo.FetishModel{}, &repo.ProductFetishModel{}, &repo.UserFetishModel{}, &repo.LikeModel{}, &repo.RecommendationModel{}, &repo.NotificationModel{}, &repo.ReviewModel{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	ctx := context.Background()

	// create category
	cat := &repo.CategoryModel{Name: "Dildos"}
	if err := db.WithContext(ctx).Create(cat).Error; err != nil {
		t.Fatalf("create category: %v", err)
	}

	// create products
	p1 := &repo.ProductModel{Name: "Soft Bunny", Description: "Plush vibrator", Price: 29.99, CategoryID: cat.ID}
	p2 := &repo.ProductModel{Name: "Latex Tail", Description: "Tail with plug", Price: 49.5, CategoryID: cat.ID}
	if err := db.WithContext(ctx).Create(p1).Error; err != nil {
		t.Fatalf("create p1: %v", err)
	}
	if err := db.WithContext(ctx).Create(p2).Error; err != nil {
		t.Fatalf("create p2: %v", err)
	}

	// create users
	u1 := &repo.UserModel{Username: "furry1", Email: "f1@example.test"}
	u2 := &repo.UserModel{Username: "furry2", Email: "f2@example.test"}
	if err := db.WithContext(ctx).Create(u1).Error; err != nil {
		t.Fatalf("create u1: %v", err)
	}
	if err := db.WithContext(ctx).Create(u2).Error; err != nil {
		t.Fatalf("create u2: %v", err)
	}

	// create fetishes
	fet1 := &repo.FetishModel{Name: "Tailplay"}
	fet2 := &repo.FetishModel{Name: "Latex"}
	if err := db.WithContext(ctx).Create(fet1).Error; err != nil {
		t.Fatalf("create fet1: %v", err)
	}
	if err := db.WithContext(ctx).Create(fet2).Error; err != nil {
		t.Fatalf("create fet2: %v", err)
	}

	// product-fetish links
	pf1 := &repo.ProductFetishModel{ProductID: p1.ID, FetishID: fet1.ID}
	pf2 := &repo.ProductFetishModel{ProductID: p2.ID, FetishID: fet2.ID}
	if err := db.WithContext(ctx).Create(pf1).Error; err != nil {
		t.Fatalf("create pf1: %v", err)
	}
	if err := db.WithContext(ctx).Create(pf2).Error; err != nil {
		t.Fatalf("create pf2: %v", err)
	}

	// likes
	l1 := &repo.LikeModel{UserID: u1.ID, ProductID: p1.ID, CreatedAt: time.Now()}
	l2 := &repo.LikeModel{UserID: u2.ID, ProductID: p1.ID, CreatedAt: time.Now()}
	if err := db.WithContext(ctx).Create(l1).Error; err != nil {
		t.Fatalf("create like1: %v", err)
	}
	if err := db.WithContext(ctx).Create(l2).Error; err != nil {
		t.Fatalf("create like2: %v", err)
	}

	// recommendations
	r1 := &repo.RecommendationModel{UserID: u1.ID, ProductID: p2.ID, Score: 0.9, CreatedAt: time.Now()}
	if err := db.WithContext(ctx).Create(r1).Error; err != nil {
		t.Fatalf("create rec1: %v", err)
	}

	// notifications
	n1 := &repo.NotificationModel{UserID: u1.ID, Title: "Sale", Body: "10% off", Read: false, CreatedAt: time.Now()}
	if err := db.WithContext(ctx).Create(n1).Error; err != nil {
		t.Fatalf("create notify1: %v", err)
	}

	// reviews
	rv1 := &repo.ReviewModel{UserID: u2.ID, ProductID: p1.ID, Rating: 5, Text: "Great!", CreatedAt: time.Now()}
	if err := db.WithContext(ctx).Create(rv1).Error; err != nil {
		t.Fatalf("create review1: %v", err)
	}

	// simple verifications
	var pc int64
	t.Run("counts", func(t *testing.T) {
		if err := db.WithContext(ctx).Model(&repo.ProductModel{}).Count(&pc).Error; err != nil {
			t.Fatalf("count products: %v", err)
		}
		if pc < 2 {
			t.Fatalf("expected >=2 products, got %d", pc)
		}
	})
}
