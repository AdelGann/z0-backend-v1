package config

import (
	"fmt"
	"log"

	"github.com/AdelGann/z0-backend-v1/models"
	"gorm.io/gorm"
)

func MigrateDB(DB *gorm.DB) {
	fmt.Println("Migrating...")

	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Org{},
		&models.Employees{},
		&models.Client{},
		&models.ClientFeedback{},
		&models.Products{},
		&models.Metrics{},
		&models.Debts{},
		&models.Income{},
		&models.DebtType{},
		&models.IncomeType{},
	}

	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Fatalf("Error al habilitar la extensión UUID: %v", err)
	}

	if err := DB.AutoMigrate(modelsToMigrate...); err != nil {
		log.Fatalf("Error en la migración: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}
