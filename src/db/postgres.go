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
	createDatabase(cfg)
	dsn := "host=" + cfg.Host + " user=" + cfg.User + " password=" + cfg.Password + " dbname=" + cfg.Name + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected")

	err = enableUUIDExtension(DB)
	if err != nil {
		log.Fatalf("Failed to enable uuid-ossp extension: %v", err)
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func enableUUIDExtension(db *gorm.DB) error {
	return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
}

func createDatabase(cfg *config.PostgresConfig) {
	dsn := "host=" + cfg.Host + " user=" + cfg.User + " password=" + cfg.Password + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode + " dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to default database: %v", err)
	}

	var dbName string
	db.Raw("SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower(?)", cfg.Name).Scan(&dbName)

	if dbName == "" {
		err = db.Exec("CREATE DATABASE " + cfg.Name).Error
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		log.Println("Database created")
	} else {
		log.Println("Database already exists")
	}
}
