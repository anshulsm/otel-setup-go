package controllers

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"gocrudapp/config"
	"gocrudapp/models"
	"gocrudapp/services"
	"net/http"
)

var tracer = otel.Tracer("go-crud-app/controllers/product-controller")
var productService = services.NewProductService(config.DB)

// GetProducts handles GET /products
func GetProducts(c *gin.Context) {
	ctx, span := tracer.Start(c, "GetProducts")
	defer span.End()

	products, err := productService.GetAllProducts(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct handles GET /products/:id
func GetProduct(c *gin.Context) {
	ctx, span := tracer.Start(c, "GetProductById")
	defer span.End()

	product, err := productService.GetProductByID(ctx, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct handles POST /products
func CreateProduct(c *gin.Context) {
	ctx, span := tracer.Start(c, "CreateProduct")
	defer span.End()

	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := productService.CreateProduct(ctx, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdProduct)
}

// UpdateProduct handles PUT /products/:id
func UpdateProduct(c *gin.Context) {
	ctx, span := tracer.Start(c, "UpdateProduct")
	defer span.End()

	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := productService.UpdateProduct(ctx, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct handles DELETE /products/:id
func DeleteProduct(c *gin.Context) {
	ctx, span := tracer.Start(c, "DeleteProduct")
	defer span.End()

	err := productService.DeleteProduct(ctx, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
