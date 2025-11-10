package router

import (
	"tctApi/internal/auth"
	usrHandler "tctApi/internal/user/handler"
	usrRepo "tctApi/internal/user/repository"
	usrUsecase "tctApi/internal/user/usecase"

	orgHandler "tctApi/internal/organizer/handler"
	orgRepo "tctApi/internal/organizer/repository"
	orgUsecase "tctApi/internal/organizer/usecase"

	rlHandler "tctApi/internal/role/handler"
	rlRepo "tctApi/internal/role/repository"
	rlUsecase "tctApi/internal/role/usecase"

	ttHandler "tctApi/internal/ticket-tier/handler"
	ttRepo "tctApi/internal/ticket-tier/repository"
	ttUsecase "tctApi/internal/ticket-tier/usecase"

	tHandler "tctApi/internal/ticket/handler"
	tRepo "tctApi/internal/ticket/repository"
	tUsecase "tctApi/internal/ticket/usecase"

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

	roleRepo := rlRepo.NewRoleRepository(db)
	roleUsecase := rlUsecase.NewRoleUsecase(roleRepo)
	roleHandler := rlHandler.NewRoleHandler(roleUsecase)

	ticketTierRepo := ttRepo.NewTicketTierRepository(db)
	ticketTierUsecase := ttUsecase.NewTicketTierUsecase(ticketTierRepo)
	ticketTierHandler := ttHandler.NewTicketTierHandler(ticketTierUsecase)

	ticketRepo := tRepo.NewTicketRepository(db)
	ticketUsecase := tUsecase.NewTicketUsecase(ticketRepo, ticketTierRepo)
	ticketHandler := tHandler.NewTicketHandler(ticketUsecase)

	CorsConfig(r)
	api := r.Group("/api/v1")
	usersGroup := api.Group("/users")
	usersGroup.POST("/register", userHandler.Register)
	usersGroup.POST("/login", userHandler.Login)

	authGroup := api.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/refresh", authHandler.Refresh)

	organizerGroup := api.Group("/organizers", auth.AuthMiddleware(jwtService))
	organizerGroup.POST("/create", organizerHandler.Create)
	organizerGroup.GET("/all", organizerHandler.FindAll)
	organizerGroup.GET("/:id", organizerHandler.FindById)
	organizerGroup.PUT("/update/:id", organizerHandler.Update)
	organizerGroup.DELETE("/delete/:id", organizerHandler.Delete)

	roleGroup := api.Group("/roles", auth.AuthMiddleware(jwtService))
	roleGroup.POST("create", roleHandler.Create)
	roleGroup.GET("/all", roleHandler.FindAll)
	roleGroup.GET("/:id", roleHandler.FindById)
	roleGroup.PUT("/update/:id", roleHandler.Update)
	roleGroup.DELETE("/delete/:id", roleHandler.Delete)

	ticketTierGroup := api.Group("/ticket-tiers", auth.AuthMiddleware(jwtService))
	ticketTierGroup.POST("/create", ticketTierHandler.Create)
	ticketTierGroup.GET("/all", ticketTierHandler.FindAll)
	ticketTierGroup.GET("/:id", ticketTierHandler.FindById)
	ticketTierGroup.PUT("/update/:id", ticketTierHandler.Update)
	ticketTierGroup.DELETE("/delete/:id", ticketTierHandler.Delete)

	ticketGroup := api.Group("/tickets", auth.AuthMiddleware(jwtService))
	ticketGroup.POST("/create", ticketHandler.Create)
	ticketGroup.GET("/all", ticketHandler.FindAll)
	ticketGroup.GET("/:id", ticketHandler.FindById)
	ticketGroup.PUT("/update/:id", ticketHandler.Update)
	ticketGroup.DELETE("/delete/:id", ticketHandler.Delete)

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
