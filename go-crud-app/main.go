// main.go
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"gocrudapp/config"
	"gocrudapp/models"
	"gocrudapp/routes"
	"log"
)

func main() {
	// tracer setup
	traceProvider, err := config.StartTracing()
	if err != nil {
		log.Fatalf("traceprovider: %v", err)
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Fatalf("traceprovider: %v", err)
		}
	}()

	_ = traceProvider.Tracer("go-crud-app")

	// Initialize database
	config.InitDB()

	// Run AutoMigrate to keep DB schema up to date
	mgErr := config.DB.AutoMigrate(&models.Product{})
	if mgErr != nil {
		log.Fatal("Error while auto migrating: ", err)
		return
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.ProductRoutes(router)

	// Start server
	err = router.Run(":8089")
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
