package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	"warehouse-management-system-eFishery/repository"
)

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
