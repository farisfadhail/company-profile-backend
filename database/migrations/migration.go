package migrations

import (
	"fmt"
	"log"
	"plastindo-back-end/database"
)

func RunMigration() {
	db := database.DatabaseInit()

	err := db.AutoMigrate()

	if err != nil {
		log.Println(err)
	}
	
	fmt.Println("Database Migrated")
}