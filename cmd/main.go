package main

import (
	"fmt"
	"os"
	"tctApi/internal/router"
	"tctApi/pkg/config"
	"tctApi/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		database.RunMigration(db)
		return
	}

	gin.SetMode(cfg.AppMode)
	r := gin.Default()
	router.SetupRoutes(r, db, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
