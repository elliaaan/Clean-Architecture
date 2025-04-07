package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	// Create a Gin router instance
	r := gin.Default()

	// Initialize Resty HTTP client for forwarding requests
	client := resty.New()

	// 🔽 INVENTORY SERVICE ROUTING (localhost:8080)
	inventoryURL := "http://localhost:8080"

	// GET /products → forwards to inventory-service
	r.GET("/products", func(c *gin.Context) {
		forwardGet(c, client, inventoryURL+"/products")
	})

	// 🔽 ORDER SERVICE ROUTING (localhost:8081)
	orderURL := "http://localhost:8081"

	// POST /orders → create new order
	r.POST("/orders", func(c *gin.Context) {
		forwardBody(c, client, orderURL+"/orders", http.MethodPost)
	})

	// GET /orders?user_id=1 → get orders by user
	r.GET("/orders", func(c *gin.Context) {
		userID := c.Query("user_id")
		url := orderURL + "/orders"
		if userID != "" {
			url += "?user_id=" + userID
		}
		forwardGet(c, client, url)
	})

	// GET /orders/:id → get order by ID
	r.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		forwardGet(c, client, orderURL+"/orders/"+id)
	})

	// PATCH /orders/:id → update order status
	r.PATCH("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		forwardBody(c, client, orderURL+"/orders/"+id, http.MethodPatch)
	})

	// Start the API Gateway on port 8090
	r.Run(":8090")
}

// Helper function to forward GET requests to microservices
func forwardGet(c *gin.Context, client *resty.Client, url string) {
	resp, err := client.R().Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

// Helper function to forward POST and PATCH requests with body
func forwardBody(c *gin.Context, client *resty.Client, url string, method string) {
	body, _ := io.ReadAll(c.Request.Body)
	req := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body)

	var resp *resty.Response
	var err error

	switch method {
	case http.MethodPost:
		resp, err = req.Post(url)
	case http.MethodPatch:
		resp, err = req.Patch(url)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}
