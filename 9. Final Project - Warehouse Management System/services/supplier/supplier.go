package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/supplier"
	repository "warehouse-management-system-eFishery/repository/supplier"
)

type InterfaceSupplierService interface {
	CreateSupplier(supplier entity.Supplier) (entity.Supplier, error)
	GetAllSupplier() ([]entity.Supplier, error)
	GetSupplierByID(id int) (entity.Supplier, error)
	UpdateSupplier(id int, supplierReq entity.UpdateSupplier) (entity.Supplier, error)
	DeleteSupplier(id int) error
}

type SupplierService struct {
	supplierRepository repository.InterfaceSupplierRepository
}

func NewSupplierService(supplierRepository repository.InterfaceSupplierRepository) *SupplierService {
	return &SupplierService{supplierRepository}
}

func (service SupplierService) CreateSupplier(supplier entity.Supplier) (entity.Supplier, error) {
	sup := entity.Supplier{
		SupplierName: supplier.SupplierName,
		SupplierDesc: supplier.SupplierDesc,
		Phone:        supplier.Phone,
		Email:        supplier.Email,
		CreatedAt:    time.Now().Format(time.RFC3339Nano),
	}
	supplier, err := service.supplierRepository.Store(sup)
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (service SupplierService) GetAllSupplier() ([]entity.Supplier, error) {
	supplier, err := service.supplierRepository.FindAll()
	if err != nil {
		return supplier, nil
	}
	return supplier, nil
}

func (service SupplierService) GetSupplierByID(id int) (entity.Supplier, error) {
	supplier, err := service.supplierRepository.FindByID(id)
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}

func (service SupplierService) UpdateSupplier(id int, supplierReq entity.UpdateSupplier) (entity.Supplier, error) {
	supplier, err := service.supplierRepository.FindByID(id)
	if err != nil {
		return entity.Supplier{}, err
	}
	s := entity.Supplier{
		ID:           supplier.ID,
		SupplierName: supplierReq.SupplierName,
		SupplierDesc: supplierReq.SupplierDesc,
		Phone:        supplierReq.Phone,
		Email:        supplierReq.Email,
		CreatedAt:    supplier.CreatedAt,
		UpdatedAt:    time.Now().Format(time.RFC3339Nano),
	}
	supplier, err = service.supplierRepository.Update(s)
	if err != nil {
		return entity.Supplier{}, err
	}
	return supplier, nil

}

func (service SupplierService) DeleteSupplier(id int) error {
	err := service.supplierRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
