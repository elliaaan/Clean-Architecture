package main

import (
	"api-gateway/internal/inventory"
	"api-gateway/internal/order"
	"log"
	"net/http"
	"strconv"

	orderpb "github.com/elliaaan/proto-gen/pb/order/github.com/elliaaan/proto-gen/pb/order"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gRPC-клиенты
	invClient := inventory.NewInventoryClient("localhost:8080")
	orderClient := order.NewOrderClient("localhost:8081")

	r.GET("/products", func(c *gin.Context) {
		products, err := invClient.ListProducts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	r.POST("/orders", func(c *gin.Context) {
		var req struct {
			UserID     uint64  `json:"user_id"`
			ProductID  uint64  `json:"product_id"`
			Quantity   uint32  `json:"quantity"`
			TotalPrice float64 `json:"total_price"`
			Status     string  `json:"status"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		order := &orderpb.Order{
			UserId:     req.UserID,
			ProductId:  req.ProductID,
			Quantity:   req.Quantity,
			TotalPrice: req.TotalPrice,
			Status:     req.Status,
		}

		res, err := orderClient.CreateOrder(order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res.Order)
	})

	r.GET("/orders", func(c *gin.Context) {
		res, err := orderClient.ListOrders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.Orders)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		res, err := orderClient.GetOrderByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.Order)
	})

	r.PUT("/orders/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		var req struct {
			UserID     uint64  `json:"user_id"`
			ProductID  uint64  `json:"product_id"`
			Quantity   uint32  `json:"quantity"`
			TotalPrice float64 `json:"total_price"`
			Status     string  `json:"status"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		order := &orderpb.Order{
			Id:         id,
			UserId:     req.UserID,
			ProductId:  req.ProductID,
			Quantity:   req.Quantity,
			TotalPrice: req.TotalPrice,
			Status:     req.Status,
		}

		res, err := orderClient.UpdateOrder(order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res.Order)
	})

	r.DELETE("/orders/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		_, err = orderClient.DeleteOrder(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
	})

	log.Println("API Gateway is running on port 8090")
	r.Run(":8090")
}
