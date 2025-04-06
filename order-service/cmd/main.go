package main

import (
	"order-service/db"
	"order-service/internal/order"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	repo := &order.Repository{DB: database}
	service := &order.Service{Repo: repo}
	handler := &order.Handler{Service: service}

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":8081") // Order Service работает на порту 8081
}
