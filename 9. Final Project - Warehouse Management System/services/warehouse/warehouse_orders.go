package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	repository "warehouse-management-system-eFishery/repository/warehouse"
)

type InterfaceWarehouseOrderService interface {
	CreateWarehouseOrder(entity.CreateWarehouseOrders) (entity.WarehouseOrders, error)
	FindAllWarehouseOrder() ([]entity.WarehouseOrders, error)
	FindWarehouseOrderById(int) (entity.WarehouseOrders, error)
	UpdateWarehouseOrderById(int, entity.UpdateWarehouseOrders) (entity.WarehouseOrders, error)
	FindWarehouseOrderByProductStatus(productStatus string) ([]entity.WarehouseOrders, error)
	FindWarehouseOrderByProductMark(productMark string) ([]entity.WarehouseOrders, error)
	DeleteWarehouseOrderById(int) error
}
type WarehouseOrderService struct {
	warehouseOrderRepository repository.InterfaceWarehouseOrderRepository
}

func NewWarehouseOrderService(warehouseOrderRepository repository.InterfaceWarehouseOrderRepository) *WarehouseOrderService {
	return &WarehouseOrderService{warehouseOrderRepository}
}

// // Storing warehouse order data to database
func (service WarehouseOrderService) CreateWarehouseOrder(warehouseReq entity.CreateWarehouseOrders) (entity.WarehouseOrders, error) {
	wos := entity.WarehouseOrders{
		WorkerUUID:    warehouseReq.WorkerUUID,
		WarehouseId:   warehouseReq.WarehouseId,
		OrderId:       warehouseReq.OrderId,
		ProductStatus: warehouseReq.ProductStatus,
		ProductMark:   warehouseReq.ProductMark,
		CreatedAt:     time.Now().Format(time.RFC3339Nano),
	}
	warehouseOrder, err := service.warehouseOrderRepository.Store(wos)
	if err != nil {
		return warehouseOrder, err
	}
	return warehouseOrder, nil
}

// // Find all warehouse order data from database
func (service WarehouseOrderService) FindAllWarehouseOrder() ([]entity.WarehouseOrders, error) {
	warehouseOrders, err := service.warehouseOrderRepository.FindAll()
	if err != nil {
		return warehouseOrders, err
	}
	return warehouseOrders, nil
}

// Find warehouse order data by id from database
func (service WarehouseOrderService) FindWarehouseOrderById(id int) (entity.WarehouseOrders, error) {
	warehouseOrder, err := service.warehouseOrderRepository.FindByID(id)
	if err != nil {
		return warehouseOrder, err
	}
	return warehouseOrder, nil
}

// Find warehouse order data by product status from database
func (service WarehouseOrderService) FindWarehouseOrderByProductStatus(productStatus string) ([]entity.WarehouseOrders, error) {
	warehouseOrder, err := service.warehouseOrderRepository.FindByProductStatus(productStatus)
	if err != nil {
		return warehouseOrder, err
	}
	return warehouseOrder, nil
}

// Find warehouse order data by product mark from database
func (service WarehouseOrderService) FindWarehouseOrderByProductMark(productMark string) ([]entity.WarehouseOrders, error) {
	warehouseOrder, err := service.warehouseOrderRepository.FindByProductMark(productMark)
	if err != nil {
		return warehouseOrder, err
	}
	return warehouseOrder, nil
}

// Update warehouse order data by id from database
func (service WarehouseOrderService) UpdateWarehouseOrderById(id int, warehouseReq entity.UpdateWarehouseOrders) (entity.WarehouseOrders, error) {
	wos := entity.WarehouseOrders{
		WorkerUUID:    warehouseReq.WorkerUUID,
		WarehouseId:   warehouseReq.WarehouseId,
		OrderId:       warehouseReq.OrderId,
		ProductStatus: warehouseReq.ProductStatus,
		ProductMark:   warehouseReq.ProductMark,
		UpdatedAt:     time.Now().Format(time.RFC3339Nano),
	}
	warehouseOrder, err := service.warehouseOrderRepository.UpdateByID(id, wos)
	if err != nil {
		return warehouseOrder, err
	}
	return warehouseOrder, nil
}

// Delete warehouse order data by id from database
func (service WarehouseOrderService) DeleteWarehouseOrderById(id int) error {
	err := service.warehouseOrderRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
