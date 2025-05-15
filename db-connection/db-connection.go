package dbconnection

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Глобална променлива
var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	// Зареждаме .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Запазваме в глобалната
	DB = db

	// // AutoMigrate
	// err = db.AutoMigrate(&databasetypes.Movie{}, &databasetypes.Actor{}, &databasetypes.MovieImage{})
	// if err != nil {
	// 	log.Fatal("Migration failed:", err)
	// }

	fmt.Println("✅ Successfully connected and migrated the PostgreSQL database.")
	return db, nil
}
