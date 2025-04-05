package product

import (
	"inventory-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/products", h.CreateProduct)
	r.GET("/products", h.GetProducts)
	r.PATCH("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	r.GET("/products/:id", h.GetProductByID)

}

func (h *Handler) CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func (h *Handler) GetProducts(c *gin.Context) {
	products, err := h.Service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.Service.UpdateProduct(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func parseUintParam(s string, out *uint) error {
	var temp uint64
	var err error
	if temp, err = strconv.ParseUint(s, 10, 32); err != nil {
		return err
	}
	*out = uint(temp)
	return nil
}
func (h *Handler) DeleteProduct(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.Service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
func (h *Handler) GetProductByID(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.Service.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
