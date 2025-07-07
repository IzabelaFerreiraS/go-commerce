package routes

import (
	"github.com/gin-gonic/gin"

	"go-commerce/handler"
	"go-commerce/services"
	"go-commerce/repositories"
	"go-commerce/config"
)

func Initialize() {
	router := gin.Default()

	db := config.GetSQLite()

	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	saleRepo := repositories.NewSaleRepository(db)

	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	saleService := services.NewSaleService(saleRepo, productRepo)

	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	saleHandler := handler.NewSaleHandler(saleService)

	api := router.Group("/api")

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUserHandler)
		userRoutes.GET("/", userHandler.ListUserHandler)
		userRoutes.GET("/:id", userHandler.ShowUserHandler)
		userRoutes.PUT("/:id", userHandler.UpdateUserHandler)
		userRoutes.DELETE("/:id", userHandler.DeleteUserHandler)
	}

	productRoutes := api.Group("/products")
	{
		productRoutes.POST("/", productHandler.CreateProductHandler)
		productRoutes.GET("/", productHandler.ListProductHandler)
		productRoutes.GET("/:id", productHandler.ShowProductHandler)
		productRoutes.PUT("/:id", productHandler.UpdateProductHandler)
		productRoutes.DELETE("/:id", productHandler.DeleteProductHandler)
	}

	saleRoutes := api.Group("/sales")
	{
		saleRoutes.POST("/", saleHandler.CreateSaleHandler)
		saleRoutes.GET("/", saleHandler.ListSalesHandler)
		saleRoutes.GET("/:id", saleHandler.ShowSaleHandler)
	}

	router.Run(":8080")
}
