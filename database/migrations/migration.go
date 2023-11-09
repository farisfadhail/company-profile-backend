package migrations

import (
	"fmt"
	"log"
	"plastindo-back-end/database"
	"plastindo-back-end/models/entity"
)

func RunMigration() {
	db := database.DatabaseInit()

	// db.Migrator().DropTable(&entity.User{}, &entity.ParentCategory{}, &entity.ProductCategory{}, &entity.Product{}, &entity.ImageGallery{})
	err := db.AutoMigrate(&entity.User{}, &entity.ParentCategory{}, &entity.ProductCategory{}, &entity.Product{}, &entity.ImageGallery{})

	if err != nil {
		log.Println(err)
	}
	
	fmt.Println("Database Migrated")
}