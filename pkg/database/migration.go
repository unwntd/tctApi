package database

import (
	"log"
	"tctApi/internal/auth"
	"tctApi/internal/organizer"
	"tctApi/internal/role"
	"tctApi/internal/ticket"
	tickettier "tctApi/internal/ticket-tier"
	users "tctApi/internal/user"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	AutoMigratableModels := []interface{}{
		&organizer.Organizer{},
		&users.User{},
		&auth.RefreshToken{},
		&role.Role{},
		&tickettier.TicketTier{},
		&ticket.Ticket{},
	}

	err := db.AutoMigrate(AutoMigratableModels...)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully")
}
