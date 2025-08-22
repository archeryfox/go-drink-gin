package main

import (
	"fmt"
	"log"
	"os"

	"go-gin-hello/internal/cache"
	"go-gin-hello/internal/db"
	repo "go-gin-hello/internal/repository"
	service "go-gin-hello/internal/service"
	"go-gin-hello/router"
)

// @title FurryShop API
// @version 1.0
// @description Пример API-магазина (демо). Использует GORM, Redis, DDD-слои.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=1234 dbname=furryshop port=5432 sslmode=disable"
	}

	gdb, err := db.NewGorm(dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate models
	if err := gdb.AutoMigrate(
		&repo.UserModel{},
		&repo.CategoryModel{},
		&repo.ProductModel{},
		&repo.FetishModel{},
		&repo.ProductFetishModel{},
		&repo.UserFetishModel{},
		&repo.LikeModel{},
		&repo.RecommendationModel{},
		&repo.NotificationModel{},
		&repo.ReviewModel{},
	); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	// Redis
	redisClient, err := cache.NewRedisClient()
	if err != nil {
		log.Printf("redis not available: %v, continuing without cache", err)
	}

	// Repos and services
	productRepo := repo.NewGormProductRepository(gdb)
	fetishRepo := repo.NewGormFetishRepository(gdb)
	likeRepo := repo.NewGormLikeRepository(gdb)
	notificationRepo := repo.NewGormNotificationRepository(gdb)
	recommendationRepo := repo.NewGormRecommendationRepository(gdb)
	productSvc := service.NewProductService(productRepo, redisClient)

	r := router.SetupRouter(productSvc, fetishRepo, likeRepo, notificationRepo, recommendationRepo)
	addr := ":8080"
	fmt.Printf("listening on %s\n", addr)
	err = r.Run(addr)
	if err != nil {
		return
	}
}
