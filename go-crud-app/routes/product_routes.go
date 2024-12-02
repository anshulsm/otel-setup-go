// routes/product_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"gocrudapp/controllers"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProduct)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
}
