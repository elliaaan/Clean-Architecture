package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	r := gin.Default()
	client := resty.New()

	// INVENTORY SERVICE (localhost:8080)
	inventoryURL := "http://localhost:8080"

	r.GET("/products", func(c *gin.Context) {
		forwardGet(c, client, inventoryURL+"/products")
	})

	// ORDER SERVICE (localhost:8081)
	orderURL := "http://localhost:8081"

	r.POST("/orders", func(c *gin.Context) {
		forwardBody(c, client, orderURL+"/orders", http.MethodPost)
	})

	r.GET("/orders", func(c *gin.Context) {
		userID := c.Query("user_id")
		url := orderURL + "/orders"
		if userID != "" {
			url += "?user_id=" + userID
		}
		forwardGet(c, client, url)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		forwardGet(c, client, orderURL+"/orders/"+id)
	})

	r.PATCH("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		forwardBody(c, client, orderURL+"/orders/"+id, http.MethodPatch)
	})

	r.Run(":8090") // API Gateway on port 8090
}

func forwardGet(c *gin.Context, client *resty.Client, url string) {
	resp, err := client.R().Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

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
