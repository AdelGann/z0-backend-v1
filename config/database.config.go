package config

import (
	"fmt"
	"log"
	"os"

	"github.com/AdelGann/z0-backend-v1/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// load .env variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error charging file .env")
	}

	// setting up postgres config
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("ERR connecting to database:", err)
	}

	DB = db

	MigrateDB()

	fmt.Println("Database connected and migrated successfully.")

}
func MigrateDB() {
	// formatter
	fmt.Println("===============================================")
	fmt.Println("Migrating...")


	modelsToMigrate := []interface{}{
		&models.User{},
		// add more there
	}
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	DB.AutoMigrate(modelsToMigrate...)

	fmt.Println("Migration completed successfully.")
}
