package api

import (
	"base-trade-rest/api/handler"
	"base-trade-rest/core/repository"
	"base-trade-rest/core/service"
	"base-trade-rest/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db := database.GetDB()

	// Register Handler
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	router := gin.Default()

	user := router.Group("auth")
	{
		user.POST("/register", authHandler.UserRegister)
		user.POST("/login", authHandler.UserLogin)
	}

	return router
}
