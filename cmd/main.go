package main

import (
	"fmt"
	"tctApi/internal/router"
	"tctApi/pkg/config"
	"tctApi/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)

	r := gin.Default()
	router.SetupRoutes(r, db, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
