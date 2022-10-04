package services

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/repository"
)

type InterfaceWarehouseCategoryService interface {
	CreateWarehouseCategory(warehouseCategory entity.WarehouseCategories) (entity.WarehouseCategories, error)
	GetAllWarehouseCategory() ([]entity.WarehouseCategories, error)
	GetCategoryByID(id int) (entity.WarehouseCategories, error)
	UpdateWarehouseCategory(id int, warehouseCategories entity.WarehouseCategories) (entity.WarehouseCategories, error)
	DeleteWarehouseCategory(id int) error
}

type WarehouseCategoryService struct {
	warehouseCategoryRepository repository.InterfaceWarehouseCategoriesRepository
}

func NewWarehouseCategoryService(warehouseCategoryRepository repository.InterfaceWarehouseCategoriesRepository) *WarehouseCategoryService {
	return &WarehouseCategoryService{warehouseCategoryRepository}
}

func (service WarehouseCategoryService) CreateWarehouseCategory(warehouse entity.WarehouseCategories) (entity.WarehouseCategories, error) {
	was := entity.WarehouseCategories{
		ID:           warehouse.ID,
		CategoryName: warehouse.CategoryName,
		CategoryDesc: warehouse.CategoryDesc,
	}
	warehouseCategory, err := service.warehouseCategoryRepository.Store(was)
	if err != nil {
		return warehouseCategory, err
	}
	return warehouseCategory, nil
}

func (service WarehouseCategoryService) GetAllCategory() ([]entity.WarehouseCategories, error) {
	warehouseCategory, err := service.warehouseCategoryRepository.FindAll()
	if err != nil {
		return warehouseCategory, nil
	}
	return warehouseCategory, nil
}

func (service WarehouseCategoryService) GetCategoryByID(id int) (entity.WarehouseCategories, error) {
	warehouseCategory, err := service.warehouseCategoryRepository.FindByID(id)
	if err != nil {
		return warehouseCategory, err
	}
	return warehouseCategory, nil
}

func (service WarehouseCategoryService) UpdateWarehouseCategory(id int, warehouseReq entity.WarehouseCategories) (entity.WarehouseCategories, error) {
	warehouseCategory, err := service.warehouseCategoryRepository.FindByID(id)
	if err != nil {
		return entity.WarehouseCategories{}, err
	}
	was := entity.WarehouseCategories{
		ID:           warehouseCategory.ID,
		CategoryName: warehouseReq.CategoryName,
		CategoryDesc: warehouseReq.CategoryDesc,
	}
	warehouseCategory, err = service.warehouseCategoryRepository.Update(was)
	if err != nil {
		return entity.WarehouseCategories{}, err
	}
	return warehouseCategory, nil
}

func (service WarehouseCategoryService) DeleteWarehouseCategory(id int) error {
	err := service.warehouseCategoryRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
