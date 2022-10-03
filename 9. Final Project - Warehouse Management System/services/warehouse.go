package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/repository"
)

type InterfaceWarehouseService interface {
	CreateWarehouse(warehouse entity.Warehouse) (entity.Warehouse, error)
	GetAllWarehouse() ([]entity.Warehouse, error)
	GetWarehouseByID(id int) (entity.Warehouse, error)
	UpdateWarehouse(warehouse entity.Warehouse, id int) (entity.Warehouse, error)
	DeleteWarehouse(id int) error
}

type WarehouseService struct {
	warehouseRepository repository.InterfaceWarehouseRepository
}

func NewWarehouseService(warehouseRepository repository.InterfaceWarehouseRepository) *WarehouseService {
	return &WarehouseService{warehouseRepository}
}

func (service WarehouseService) CreateWarehouse(warehouse entity.Warehouse) (entity.Warehouse, error) {
	w := entity.Warehouse{
		WarehouseName: warehouse.WarehouseName,
		WarehouseDesc: warehouse.WarehouseDesc,
		CreatedAt:     time.Now().Format(time.RFC3339Nano),
	}
	warehouse, err := service.warehouseRepository.Store(w)
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

func (service WarehouseService) GetWarehouseByID(id int) (entity.Warehouse, error) {
	warehouse, err := service.warehouseRepository.FindByID(id)
	if err != nil {
		return warehouse, err
	}
	return warehouse, nil
}

func (service WarehouseService) UpdateWarehouse(warehouseReq entity.UpdateWarehouse, id int) (entity.WarehouseResponse, error) {
	warehouse, err := service.warehouseRepository.FindByID(id)
	if err != nil {
		return entity.WarehouseResponse{}, err
	}
	warehouseData := entity.Warehouse{
		ID:            warehouse.ID,
		WarehouseName: warehouseReq.WarehouseName,
		WarehouseDesc: warehouseReq.WarehouseDesc,
		CreatedAt:     warehouse.CreatedAt,
		UpdatedAt:     time.Now().Format(time.RFC3339Nano),
	}
	warehouse, err = service.warehouseRepository.Update(warehouseData)
	if err != nil {
		return entity.WarehouseResponse{}, err
	}
	warehouseResponse := entity.WarehouseResponse{
		WarehouseName: warehouse.WarehouseName,
		WarehouseDesc: warehouse.WarehouseDesc,
	}
	return warehouseResponse, nil
}

func (service WarehouseService) DeleteWarehouse(id int) error {
	err := service.warehouseRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
