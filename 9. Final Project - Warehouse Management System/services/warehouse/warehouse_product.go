package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	repository "warehouse-management-system-eFishery/repository/warehouse"
)

type InterfaceWarehouseProductService interface {
	CreateWarehouseProduct(entity.WarehouseProducts) (entity.WarehouseProducts, error)
	GetAllWarehouseProduct() ([]entity.WarehouseProducts, error)
	GetWarehouseProductByID(int) (entity.WarehouseProducts, error)
	UpdateWarehouseProduct(id int, warehouseReq entity.UpdateWarehouseProducts) (entity.WarehouseProductsResponse, error)
	DeleteWarehouseProduct(id int) error
}

type WarehouseProductService struct {
	warehouseProductRepository repository.InterfaceWarehouseProductRepository
}

func NewWarehouseProductService(warehouseProductRepository repository.InterfaceWarehouseProductRepository) *WarehouseProductService {
	return &WarehouseProductService{warehouseProductRepository}
}

func (service WarehouseProductService) CreateWarehouseProduct(warehouseReq entity.WarehouseProducts) (entity.WarehouseProducts, error) {
	wps := entity.WarehouseProducts{
		SKU:            warehouseReq.SKU,
		ProductName:    warehouseReq.ProductName,
		ProductDesc:    warehouseReq.ProductDesc,
		Image:          warehouseReq.Image,
		Price:          warehouseReq.Price,
		Stock:          warehouseReq.Stock,
		Weight:         warehouseReq.Weight,
		SectionPlaceID: warehouseReq.SectionPlaceID,
		CategoryID:     warehouseReq.CategoryID,
		SupplierID:     warehouseReq.SupplierID,
		WarehouseID:    warehouseReq.WarehouseID,
		CreatedAt:      time.Now().Format(time.RFC3339Nano),
	}
	warehouseProduct, err := service.warehouseProductRepository.Store(wps)
	if err != nil {
		return warehouseProduct, err
	}
	return warehouseProduct, nil
}

func (service WarehouseProductService) GetAllWarehouseProduct() ([]entity.WarehouseProducts, error) {
	warehouseProduct, err := service.warehouseProductRepository.GetAll()
	if err != nil {
		return warehouseProduct, err
	}
	return warehouseProduct, nil
}

func (service WarehouseProductService) GetWarehouseProductByID(id int) (entity.WarehouseProducts, error) {
	warehouseProduct, err := service.warehouseProductRepository.GetbyID(id)
	if err != nil {
		return warehouseProduct, err
	}
	return warehouseProduct, nil
}

func (service WarehouseProductService) UpdateWarehouseProduct(id int, warehouseReq entity.UpdateWarehouseProducts) (entity.WarehouseProductsResponse, error) {
	warehouseProduct, err := service.warehouseProductRepository.GetbyID(id)
	if err != nil {
		return entity.WarehouseProductsResponse{}, err
	}

	wps := entity.WarehouseProducts{
		ID:             warehouseProduct.ID,
		SKU:            warehouseReq.SKU,
		ProductName:    warehouseReq.ProductName,
		ProductDesc:    warehouseReq.ProductDesc,
		Image:          warehouseReq.Image,
		Price:          warehouseReq.Price,
		Stock:          warehouseReq.Stock,
		Weight:         warehouseReq.Weight,
		SectionPlaceID: warehouseReq.SectionPlaceID,
		CategoryID:     warehouseReq.CategoryID,
		SupplierID:     warehouseReq.SupplierID,
		WarehouseID:    warehouseReq.WarehouseID,
		CreatedAt:      warehouseProduct.CreatedAt,
		UpdatedAt:      time.Now().Format(time.RFC3339Nano),
	}

	warehouseProduct, err = service.warehouseProductRepository.Update(wps)
	if err != nil {
		return entity.WarehouseProductsResponse{}, err
	}

	warehouseProductResponse := entity.WarehouseProductsResponse{
		// ID:             warehouseProduct.ID,
		SKU:            warehouseProduct.SKU,
		ProductName:    warehouseProduct.ProductName,
		ProductDesc:    warehouseProduct.ProductDesc,
		Image:          warehouseProduct.Image,
		Price:          warehouseProduct.Price,
		Stock:          warehouseProduct.Stock,
		Weight:         warehouseProduct.Weight,
		SectionPlaceID: warehouseProduct.SectionPlaceID,
		CategoryID:     warehouseProduct.CategoryID,
		SupplierID:     warehouseProduct.SupplierID,
		WarehouseID:    warehouseProduct.WarehouseID,
		CreatedAt:      warehouseProduct.CreatedAt,
		UpdatedAt:      warehouseProduct.UpdatedAt,
	}
	return warehouseProductResponse, nil

}

func (service WarehouseProductService) DeleteWarehouseProduct(id int) error {
	err := service.warehouseProductRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
