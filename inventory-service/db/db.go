package db

import (
	"inventory-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=GIprime dbname=inventory_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to PostgreSQL database")
	}

	db.AutoMigrate(&models.Product{})
	return db
}
