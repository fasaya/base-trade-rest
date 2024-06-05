package api

import (
	"base-trade-rest/api/handler"
	"base-trade-rest/api/middleware"
	"base-trade-rest/core/repository"
	"base-trade-rest/core/service"
	"base-trade-rest/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db := database.GetDB()

	// Register
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	user := router.Group("auth")
	{
		user.POST("/register", authHandler.Register)
		user.POST("/login", authHandler.Login)
	}

	productRouter := router.Group("/products")
	{
		productRouter.GET("/", productHandler.Index)
		productRouter.GET("/:productUUID", productHandler.Show)

		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", productHandler.Store)

		// productRouter.PUT("/:productUUID", middleware.ProductAuthorization(), productHandler.Update)
		productRouter.DELETE("/:productUUID", middleware.ProductAuthorization(), productHandler.Delete)
	}

	return router
}
