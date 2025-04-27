package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ProductHandler handles product service-related requests
type ProductHandler struct {
	productServiceURL string
	logger            *logrus.Logger
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(productServiceURL string, logger *logrus.Logger) *ProductHandler {
	return &ProductHandler{
		productServiceURL: productServiceURL,
		logger:            logger,
	}
}

// GetProducts gets all products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	resp, err := http.Get(h.productServiceURL + "/products")
	if err != nil {
		h.logger.WithError(err).Error("Failed to get products from product service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get products",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		h.logger.WithField("status_code", resp.StatusCode).Error("Product service returned non-OK status")
		c.JSON(resp.StatusCode, gin.H{
			"error": "Product service error",
		})
		return
	}

	var products interface{}
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		h.logger.WithError(err).Error("Failed to decode product service response")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode response",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProduct gets a product by ID
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get(h.productServiceURL + "/products/" + id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get product from product service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get product",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		h.logger.WithField("status_code", resp.StatusCode).Error("Product service returned non-OK status")
		c.JSON(resp.StatusCode, gin.H{
			"error": "Product service error",
		})
		return
	}

	var product interface{}
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		h.logger.WithError(err).Error("Failed to decode product service response")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode response",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", h.productServiceURL+"/products", c.Request.Body)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create request to product service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := client.Do(req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create product in product service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		h.logger.WithField("status_code", resp.StatusCode).Error("Product service returned non-OK status")
		c.JSON(resp.StatusCode, gin.H{
			"error": "Product service error",
		})
		return
	}

	var product interface{}
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		h.logger.WithError(err).Error("Failed to decode product service response")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode response",
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}