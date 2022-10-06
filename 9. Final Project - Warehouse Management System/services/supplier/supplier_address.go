package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/supplier"
	repository "warehouse-management-system-eFishery/repository/supplier"
)

type InterfaceSupplierAddressService interface {
}

type SupplierAddressService struct {
	supplierAddressRepository repository.InterfaceSupplierAddressRepository
}

func NewSupplierAddressService(supplierAddressRepository repository.InterfaceSupplierAddressRepository) *SupplierAddressService {
	return &SupplierAddressService{supplierAddressRepository}
}

func (service SupplierAddressService) CreateSupplierAddress(supplierAddressRequest entity.SupplierAddress) (entity.SupplierAddress, error) {
	sas := entity.SupplierAddress{
		SupplierID:  supplierAddressRequest.SupplierID,
		FullAddress: supplierAddressRequest.FullAddress,
		SubDistrict: supplierAddressRequest.SubDistrict,
		District:    supplierAddressRequest.District,
		City:        supplierAddressRequest.City,
		Province:    supplierAddressRequest.Province,
		PostalCode:  supplierAddressRequest.PostalCode,
		CreatedAt:   time.Now().Format(time.RFC3339Nano),
	}
	supplierAddress, err := service.supplierAddressRepository.Store(sas)
	if err != nil {
		return supplierAddress, err
	}
	return supplierAddress, nil
}

func (service SupplierAddressService) GetAllSupplierAddress() ([]entity.SupplierAddress, error) {
	supplierAddresses, err := service.supplierAddressRepository.FindAll()
	if err != nil {
		return supplierAddresses, err
	}
	return supplierAddresses, nil
}

func (service SupplierAddressService) GetSupplierAddressByID(id int) (entity.SupplierAddress, error) {
	supplierAddress, err := service.supplierAddressRepository.FindByID(id)
	if err != nil {
		return supplierAddress, err
	}
	return supplierAddress, nil
}

func (service SupplierAddressService) UpdateSupplierAddress(supplierAddressRequest entity.UpdateSupplierAddress, id int) (entity.SupplierAddressResponse, error) {
	supplierAddress, err := service.supplierAddressRepository.FindByID(id)
	if err != nil {
		return entity.SupplierAddressResponse{}, err
	}
	supplierAddressData := entity.SupplierAddress{
		ID:          supplierAddress.ID,
		SupplierID:  supplierAddress.SupplierID,
		FullAddress: supplierAddressRequest.FullAddress,
		SubDistrict: supplierAddressRequest.SubDistrict,
		District:    supplierAddressRequest.District,
		City:        supplierAddressRequest.City,
		Province:    supplierAddressRequest.Province,
		PostalCode:  supplierAddressRequest.PostalCode,
		CreatedAt:   supplierAddress.CreatedAt,
		UpdatedAt:   time.Now().Format(time.RFC3339Nano),
	}
	supplierAddress, err = service.supplierAddressRepository.Update(supplierAddressData)
	if err != nil {
		return entity.SupplierAddressResponse{}, err
	}
	supplierAddressResponse := entity.SupplierAddressResponse{
		ID:          supplierAddress.ID,
		SupplierID:  supplierAddress.SupplierID,
		FullAddress: supplierAddress.FullAddress,
		SubDistrict: supplierAddress.SubDistrict,
		District:    supplierAddress.District,
		City:        supplierAddress.City,
		Province:    supplierAddress.Province,
		PostalCode:  supplierAddress.PostalCode,
		CreatedAt:   supplierAddress.CreatedAt,
		UpdatedAt:   supplierAddress.UpdatedAt,
	}
	return supplierAddressResponse, nil
}

func (service SupplierAddressService) DeleteSupplierAddress(id int) error {
	err := service.supplierAddressRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
