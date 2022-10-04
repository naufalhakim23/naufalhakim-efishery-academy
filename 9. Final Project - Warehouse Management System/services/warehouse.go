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

type InterfaceWarehouseAddressService interface {
	CreateWarehouseAddress(warehouseAddress entity.WarehouseAddress) (entity.WarehouseAddress, error)
	GetAllWarehouseAddress() ([]entity.WarehouseAddress, error)
	GetWarehouseAddressByID(id int) (entity.WarehouseAddress, error)
	UpdateWarehouseAddress(warehouseAddress entity.WarehouseAddress, id int) (entity.WarehouseAddress, error)
	DeleteWarehouseAddress(id int) error
}

type WarehouseAddressService struct {
	warehouseAddressRepository repository.InterfaceWarehouseAddressRepository
}

func NewWarehouseAddressService(warehouseAddressRepository repository.InterfaceWarehouseAddressRepository) *WarehouseAddressService {
	return &WarehouseAddressService{warehouseAddressRepository}
}

func (service WarehouseAddressService) CreateWarehouseAddress(warehouse entity.WarehouseAddress) (entity.WarehouseAddress, error) {
	was := entity.WarehouseAddress{
		WarehouseID: warehouse.WarehouseID,
		FullAddress: warehouse.FullAddress,
		SubDistrict: warehouse.SubDistrict,
		City:        warehouse.City,
		Province:    warehouse.Province,
		Region:      warehouse.Region,
		PostalCode:  warehouse.PostalCode,
		CreatedAt:   time.Now().Format(time.RFC3339Nano),
	}
	warehouseAddress, err := service.warehouseAddressRepository.Store(was)
	if err != nil {
		return warehouseAddress, err
	}
	return warehouseAddress, nil
}

func (service WarehouseAddressService) GetAllAddress() ([]entity.WarehouseAddress, error) {
	warehouseAddress, err := service.warehouseAddressRepository.FindAll()
	if err != nil {
		return warehouseAddress, nil
	}
	return warehouseAddress, nil
}

func (service WarehouseAddressService) GetAddressByID(id int) (entity.WarehouseAddress, error) {
	warehouseAddress, err := service.warehouseAddressRepository.FindByID(id)
	if err != nil {
		return warehouseAddress, err
	}
	return warehouseAddress, nil
}

func (service WarehouseAddressService) UpdateWarehouseAddress(id int, warehouseReq entity.UpdateWarehouseAddress) (entity.WarehouseAddressesResponse, error) {
	warehouseAddress, err := service.warehouseAddressRepository.FindByID(id)
	if err != nil {
		return entity.WarehouseAddressesResponse{}, err
	}
	was := entity.WarehouseAddress{
		ID:          warehouseAddress.ID,
		WarehouseID: warehouseReq.WarehouseID,
		FullAddress: warehouseReq.FullAddress,
		SubDistrict: warehouseReq.SubDistrict,
		City:        warehouseReq.City,
		Province:    warehouseReq.Province,
		Region:      warehouseReq.Region,
		PostalCode:  warehouseReq.PostalCode,
		CreatedAt:   warehouseAddress.CreatedAt,
		UpdatedAt:   time.Now().Format(time.RFC3339Nano),
	}
	warehouseAddress, err = service.warehouseAddressRepository.Update(was)
	if err != nil {
		return entity.WarehouseAddressesResponse{}, err
	}
	warehouseAddressResponse := entity.WarehouseAddressesResponse{
		ID:          warehouseAddress.ID,
		FullAddress: warehouseAddress.FullAddress,
		SubDistrict: warehouseAddress.SubDistrict,
		City:        warehouseAddress.City,
		Province:    warehouseAddress.Province,
		Region:      warehouseAddress.Region,
		PostalCode:  warehouseAddress.PostalCode,
		CreatedAt:   warehouseAddress.CreatedAt,
		UpdatedAt:   warehouseAddress.UpdatedAt,
	}
	return warehouseAddressResponse, nil
}

func (service WarehouseAddressService) DeleteWarehouseAddress(id int) error {
	err := service.warehouseAddressRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

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
