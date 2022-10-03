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
		PostalCode:  warehouseReq.PostalCode,
		CreatedAt:   warehouseAddress.CreatedAt,
		UpdatedAt:   time.Now().Format(time.RFC3339Nano),
	}
	warehouseAddress, err = service.warehouseAddressRepository.Update(was)
	if err != nil {
		return entity.WarehouseAddressesResponse{}, err
	}
	warehouseAddressResponse := entity.WarehouseAddressesResponse{}
	return warehouseAddressResponse, nil
}

func (service WarehouseAddressService) DeleteWarehouseAddress(id int) error {
	err := service.warehouseAddressRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
