package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {

		log.Fatal("Error loading env file")
	}

	// build the connection string and connect to the postgres instance
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s  sslmode=require TimeZone=UTC+8",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db, nil
}
