package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	repository "warehouse-management-system-eFishery/repository/warehouse"
)

type InterfaceWarehouseSectionService interface {
	CreateWarehouseSection(warehouseRequest entity.WarehouseSection) (entity.WarehouseSection, error)
	GetAllWarehouseSection() ([]entity.WarehouseSection, error)
	GetWarehouseSectionByID(id int) (entity.WarehouseSection, error)
	UpdateWarehouseSection(id int, warehouseRequest entity.UpdateWarehouseSection) (entity.WarehouseSectionResponse, error)
	DeleteWarehouseSection(id int) error
}

type WarehouseSectionService struct {
	warehouseSectionRepository repository.InterfaceWarehouseSectionRepository
}

func NewWarehouseSectionService(warehouseSectionRepository repository.InterfaceWarehouseSectionRepository) *WarehouseSectionService {
	return &WarehouseSectionService{warehouseSectionRepository}
}

// WarehouseSectionService is a service that handles all business logic related to warehouse section

func (service WarehouseSectionService) CreateWarehouseSection(warehouseRequest entity.WarehouseSection) (entity.WarehouseSection, error) {
	wss := entity.WarehouseSection{
		WarehouseID:      warehouseRequest.WarehouseID,
		InventorySection: warehouseRequest.InventorySection,
		InventoryAisle:   warehouseRequest.InventoryAisle,
		InventoryRow:     warehouseRequest.InventoryRow,
		InventoryTier:    warehouseRequest.InventoryTier,
		CreatedAt:        time.Now().Format(time.RFC3339Nano),
	}
	warehouseSection, err := service.warehouseSectionRepository.Store(wss)
	if err != nil {
		return warehouseSection, err
	}
	return warehouseSection, nil
}

func (service WarehouseSectionService) GetAllWarehouseSection() ([]entity.WarehouseSection, error) {
	warehouseSection, err := service.warehouseSectionRepository.FindAll()
	if err != nil {
		return warehouseSection, nil
	}
	return warehouseSection, nil
}

func (service WarehouseSectionService) GetWarehouseSectionByID(id int) (entity.WarehouseSection, error) {
	warehouseSection, err := service.warehouseSectionRepository.FindByID(id)
	if err != nil {
		return warehouseSection, err
	}
	return warehouseSection, nil
}

func (service WarehouseSectionService) UpdateWarehouseSection(id int, warehouseRequest entity.UpdateWarehouseSection) (entity.WarehouseSectionResponse, error) {
	warehouseSection, err := service.warehouseSectionRepository.FindByID(id)
	if err != nil {
		return entity.WarehouseSectionResponse{}, err
	}
	wss := entity.WarehouseSection{
		ID:               warehouseSection.ID,
		WarehouseID:      warehouseRequest.WarehouseID,
		InventorySection: warehouseRequest.InventorySection,
		InventoryAisle:   warehouseRequest.InventoryAisle,
		InventoryRow:     warehouseRequest.InventoryRow,
		InventoryTier:    warehouseRequest.InventoryTier,
		CreatedAt:        warehouseSection.CreatedAt,
		UpdatedAt:        time.Now().Format(time.RFC3339Nano),
	}
	warehouseSection, err = service.warehouseSectionRepository.Update(wss)
	if err != nil {
		return entity.WarehouseSectionResponse{}, err
	}
	return entity.WarehouseSectionResponse{
		ID:               warehouseSection.ID,
		WarehouseID:      warehouseSection.WarehouseID,
		InventorySection: warehouseSection.InventorySection,
		InventoryAisle:   warehouseSection.InventoryAisle,
		InventoryRow:     warehouseSection.InventoryRow,
		InventoryTier:    warehouseSection.InventoryTier,
		CreatedAt:        warehouseSection.CreatedAt,
		UpdatedAt:        warehouseSection.UpdatedAt,
	}, nil

}

func (service WarehouseSectionService) DeleteWarehouseSection(id int) error {
	err := service.warehouseSectionRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
