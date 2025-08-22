package router

import (
	_ "go-gin-hello/docs"
	"go-gin-hello/handlers"
	repo "go-gin-hello/internal/repository"
	service "go-gin-hello/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter wires HTTP handlers and accepts service & repository dependencies
func SetupRouter(productSvc service.ProductService, fr repo.FetishRepository, lr repo.LikeRepository, nr repo.NotificationRepository, rr repo.RecommendationRepository) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api/v1")
	{
		api.GET("/hello", handlers.HelloHandler)
		api.GET("/greet/:name", handlers.GreetHandler)

		products := api.Group("/products")
		ph := handlers.NewProductHandler(productSvc)
		{
			products.GET("", ph.GetProducts)
			products.GET(":id", ph.GetProduct)
			products.POST("", ph.CreateProduct)
			products.PUT(":id", ph.UpdateProduct)
			products.DELETE(":id", ph.DeleteProduct)
		}

		// fetishes
		f := handlers.NewFetishHandler(fr)
		fet := api.Group("/fetishes")
		{
			fet.POST("", f.CreateFetish)
			fet.GET("", f.ListFetishes)
			fet.GET(":id", f.GetFetish)
		}

		// likes
		lh := handlers.NewLikeHandler(lr)
		likes := api.Group("/likes")
		{
			likes.POST("", lh.CreateLike)
			likes.GET("/product/:id/count", lh.CountByProduct)
		}

		// notifications
		nh := handlers.NewNotificationHandler(nr)
		notif := api.Group("/notifications")
		{
			notif.POST("", nh.CreateNotification)
			notif.GET("/user/:user_id", nh.ListForUser)
			notif.PUT("/:id/read", nh.MarkRead)
		}

		// recommendations
		rh := handlers.NewRecommendationHandler(rr)
		rec := api.Group("/recommendations")
		{
			rec.POST("", rh.CreateRecommendation)
			rec.GET("/user/:user_id", rh.ListForUser)
		}

		// reviews & users & categories minimal endpoints
		rhv := handlers.NewReviewHandler()
		api.POST("/reviews", rhv.CreateReview)
		api.GET("/reviews/product/:product_id", rhv.ListReviews)

		uh := handlers.NewUserHandler()
		api.POST("/users", uh.CreateUser)
		api.GET("/users/:id", uh.GetUser)

		ch := handlers.NewCategoryHandler()
		api.POST("/categories", ch.CreateCategory)
		api.GET("/categories/:id", ch.GetCategory)
	}

	return r
}
