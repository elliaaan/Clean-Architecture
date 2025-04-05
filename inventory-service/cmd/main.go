package main

import (
	"inventory-service/db"
	"inventory-service/internal/product"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	repo := &product.Repository{DB: database}
	service := &product.Service{Repo: repo}
	handler := &product.Handler{Service: service}

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":8080") // Запускает сервер на localhost:8080
}
