package service

import (
	"context"

	"github.com/StephanioCS/GO-MS-PRACTICE/internal/model"
	"github.com/StephanioCS/GO-MS-PRACTICE/internal/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// Создание нового продукта
func (s *ProductService) CreateProduct(ctx context.Context, product model.Product) error {
	_, err := s.productRepo.CreateProduct(ctx, product.Name, product.Description, product.Specs, product.Weight, product.Barcode)
	return err
}

// Получение всех продуктов
func (s *ProductService) GetProducts(ctx context.Context) ([]model.Product, error) {
	return s.productRepo.GetProducts(ctx)
}

// Получение по ID
func (s *ProductService) GetProductByID(ctx context.Context, id uuid.UUID) (model.Product, error) {
	return s.productRepo.GetProductByID(ctx, id)
}

// Обновление информации 
func (s *ProductService) UpdateProduct(ctx context.Context, product model.Product) error {
	_, err := s.productRepo.UpdateProduct(ctx, product.ID, product.Name, product.Description, product.Specs, product.Weight, product.Barcode)
	return err
}

// Удаление продукта
func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return s.productRepo.DeleteProduct(ctx, id)
}
