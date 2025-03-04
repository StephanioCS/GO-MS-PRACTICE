package service

import (
	"context"

	"github.com/StephanioCS/GO-MS-PRACTICE/internal/model"
	"github.com/StephanioCS/GO-MS-PRACTICE/internal/repository"
	"github.com/google/uuid"
)

type InventoryService struct {
	inventoryRepo *repository.InventoryRepository
	analyticsRepo *repository.AnalyticsRepository
}

func NewInventoryService(inventoryRepo *repository.InventoryRepository, analyticsRepo *repository.AnalyticsRepository) *InventoryService {
	return &InventoryService{
		inventoryRepo: inventoryRepo,
		analyticsRepo: analyticsRepo,
	}
}

func (s *InventoryService) CreateInventory(ctx context.Context, productID, warehouseID uuid.UUID, quantity int, price float64, discount float64) error {
	_, err := s.inventoryRepo.CreateInventory(ctx, productID, warehouseID, quantity, price, discount)
	return err
}

func (s *InventoryService) PurchaseProducts(ctx context.Context, productID, warehouseID uuid.UUID, quantity int) (float64, error) {
	return 0.0, nil
}

func (s *InventoryService) GetAnalyticsByWarehouse(ctx context.Context, warehouseID uuid.UUID) (model.AnalyticsData, error) {
	return s.analyticsRepo.GetAnalyticsByWarehouse(ctx, warehouseID)
}
