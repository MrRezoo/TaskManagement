package db

import (
	"github.com/MrRezoo/TaskManagement/api/models"
	"github.com/MrRezoo/TaskManagement/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.PostgresConfig) {
	dsn := "host=" + cfg.Host + " user=" + cfg.User + " password=" + cfg.Password + " dbname=" + cfg.Name + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected")
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
