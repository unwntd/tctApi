package router

import (
	"tctApi/internal/user/handler"
	"tctApi/internal/user/repository"
	"tctApi/internal/user/usecase"
	"tctApi/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// User Module
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, cfg.JWTSecret)
	userHandler := handler.NewUserHandler(userUsecase)

	api := r.Group("/api/v1")
	userHandler.RegisterRoutes(api)
}
