package database

import (
	"log"
	"tctApi/internal/auth"
	"tctApi/internal/organizer"
	users "tctApi/internal/user"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(&organizer.Organizer{}, &users.User{}, &auth.RefreshToken{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully")
}
