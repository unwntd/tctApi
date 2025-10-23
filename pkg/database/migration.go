package database

import (
	"fmt"
	"log"
	"tctApi/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(cfg *config.Config) {
	m, err := migrate.New(
		"file://pkg/database/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=require", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName),
	)
	if err != nil {
		log.Fatal("Failed to initialize migration:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migrated successfully")
}
