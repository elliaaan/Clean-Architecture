package db

import (
	"inventory-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// 1️ Connection string (DSN) to PostgreSQL
	dsn := "host=localhost user=postgres password=GIprime dbname=inventory_db port=5432 sslmode=disable"

	// 2️ Connect to PostgreSQL using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If connection fails → stop the program
		panic("failed to connect to PostgreSQL database")
	}

	// 3️ Auto-migrate the Product model (create table if not exists)
	db.AutoMigrate(&models.Product{})

	// 4️ Return DB object to be used in Repository
	return db
}
