package main

import (
	"inventory-service/db"
	"inventory-service/internal/product"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1 Initialize database connection
	database := db.InitDB()

	// 2️ Set up the Repository layer (talks to DB)
	repo := &product.Repository{DB: database}

	// 3️ Set up the Service layer (business logic)
	service := &product.Service{Repo: repo}

	// 4️ Set up the Handler layer (HTTP layer)
	handler := &product.Handler{Service: service}

	// 5️ Create Gin router
	r := gin.Default()

	// 6️ Register product routes (CRUD operations)
	handler.RegisterRoutes(r)

	// 7️ Start Inventory Service on port 8080
	r.Run(":8080")
}
