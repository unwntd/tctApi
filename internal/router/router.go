package router

import (
	"tctApi/internal/auth"
	usrHandler "tctApi/internal/user/handler"
	usrRepo "tctApi/internal/user/repository"
	usrUsecase "tctApi/internal/user/usecase"

	orgHandler "tctApi/internal/organizer/handler"
	orgRepo "tctApi/internal/organizer/repository"
	orgUsecase "tctApi/internal/organizer/usecase"

	"tctApi/pkg/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	userRepo := usrRepo.NewUserRepository(db)
	userUsecase := usrUsecase.NewUserUsecase(userRepo, cfg.JWTSecret)
	userHandler := usrHandler.NewUserHandler(userUsecase)

	refreshRepo := auth.NewRefreshTokenRepository(db)
	jwtService := auth.NewJWTService(cfg.JWTSecret, 24*time.Hour)
	authService := auth.NewAuthService(jwtService, userRepo, refreshRepo)
	authHandler := auth.NewAuthHandler(authService)

	organizerRepo := orgRepo.NewOrganizerRepository(db)
	organizerUsecase := orgUsecase.NewOrganizerUsecase(organizerRepo)
	organizerHandler := orgHandler.NewOrganizerHandler(organizerUsecase)

	CorsConfig(r)
	api := r.Group("/api/v1")
	usersGroup := api.Group("/users")
	usersGroup.POST("/register", userHandler.Register)
	usersGroup.POST("/login", userHandler.Login)

	authGroup := api.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/refresh", authHandler.Refresh)

	organizerGroup := api.Group("/organizers", auth.AuthMiddleware(jwtService))
	organizerGroup.GET("/all", organizerHandler.FindAll)
	organizerGroup.POST("/create", organizerHandler.Create)
	organizerGroup.PUT("/update/:id", organizerHandler.Update)
	organizerGroup.DELETE("/delete/:id", organizerHandler.Delete)
}

func CorsConfig(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // allow requests from any origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
