package services

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/repository"
)

type InterfaceWarehouseService interface {
	CreateWarehouse(warehouse entity.Warehouse) (entity.Warehouse, error)
	GetAllWarehouse() ([]entity.Warehouse, error)
	GetWarehouseByID(id string) (entity.Warehouse, error)
	UpdateWarehouse(warehouse entity.Warehouse) (entity.Warehouse, error)
	DeleteWarehouse(id string) (string, error)
}

type WarehouseService struct {
	warehouseRepository repository.InterfaceWarehouseRepository
}

func NewWarehouseService(warehouseRepository repository.InterfaceWarehouseRepository) *WarehouseService {
	return &WarehouseService{warehouseRepository}
}

func (service WarehouseService) CreateWarehouse(warehouse entity.Warehouse) (entity.Warehouse, error) {
	warehouse, err := service.warehouseRepository.Store(warehouse)
	if err != nil {
		return warehouse, err
	}
	return warehouse, nil
}

func (service WarehouseService) GetAllWarehouse() ([]entity.Warehouse, error) {
	warehouses, err := service.warehouseRepository.FindAll()
	if err != nil {
		return warehouses, err
	}
	return warehouses, nil
}
