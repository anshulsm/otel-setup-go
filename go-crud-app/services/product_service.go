package services

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"gocrudapp/models"
	"gorm.io/gorm"
	"log"
)

var serviceTracer = otel.Tracer("go-crud-app/services/product-service")

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	ctx, span := serviceTracer.Start(ctx, "GetAllProducts")
	defer span.End()

	var products []models.Product
	result := s.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	ctx, span := serviceTracer.Start(ctx, "GetProductByID")
	defer span.End()

	var product models.Product
	result := s.DB.First(&product, id)
	if result.Error != nil {
		return nil, errors.New("Product not found")
	}
	return &product, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	ctx, span := serviceTracer.Start(ctx, "CreateProduct")
	defer span.End()

	result := s.DB.Create(&product)
	if result.Error != nil {
		log.Println("Error occurred while creating product: ", result.Error)
		return nil, result.Error
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	ctx, span := serviceTracer.Start(ctx, "UpdateProduct")
	defer span.End()

	result := s.DB.Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	ctx, span := serviceTracer.Start(ctx, "DeleteProduct")
	defer span.End()

	var product models.Product
	result := s.DB.First(&product, id)
	if result.Error != nil {
		return errors.New("Product not found")
	}

	deleteResult := s.DB.Delete(&product)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
