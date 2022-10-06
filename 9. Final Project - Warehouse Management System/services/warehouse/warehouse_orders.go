package services

import (
	"time"
	entity "warehouse-management-system-eFishery/entity/warehouse"
	repository "warehouse-management-system-eFishery/repository/warehouse"
)

type InterfaceWarehouseOrderService interface {
	CreateWarehouseOrder(entity.CreateWarehouseOrders) (entity.WarehouseOrdersResponse, error)
	FindAllWarehouseOrder() ([]entity.WarehouseOrdersResponse, error)
	FindWarehouseOrderById(int) (entity.WarehouseOrdersResponse, error)
	UpdateWarehouseOrderById(int, entity.UpdateWarehouseOrders) (entity.WarehouseOrdersResponse, error)
	FindWarehouseOrderByProductStatus(productStatus string) ([]entity.WarehouseOrdersResponse, error)
	FindWarehouseOrderByProductMark(productMark string) ([]entity.WarehouseOrdersResponse, error)
	DeleteWarehouseOrderById(int) error
}
type WarehouseOrderService struct {
	warehouseOrderRepository repository.InterfaceWarehouseOrderRepository
}

func NewWarehouseOrderService(warehouseOrderRepository repository.InterfaceWarehouseOrderRepository) *WarehouseOrderService {
	return &WarehouseOrderService{warehouseOrderRepository}
}

// // Storing warehouse order data to database
func (service WarehouseOrderService) CreateWarehouseOrder(warehouseReq entity.CreateWarehouseOrders) (entity.WarehouseOrdersResponse, error) {
	wos := entity.WarehouseOrders{
		WorkerUUID:    warehouseReq.WorkerUUID,
		WarehouseId:   warehouseReq.WarehouseId,
		OrderId:       warehouseReq.OrderId,
		ProductStatus: warehouseReq.ProductStatus,
		ProductMark:   warehouseReq.ProductMark,
		CreatedAt:     time.Now().Format(time.RFC3339Nano),
	}
	warehouseOrder, err := service.warehouseOrderRepository.Store(wos)
	warehouseOrdersResponse := entity.WarehouseOrdersResponse{
		ID:            warehouseOrder.ID,
		WorkerUUID:    warehouseOrder.WorkerUUID,
		WarehouseId:   warehouseOrder.WarehouseId,
		Warehouse:     warehouseOrder.Warehouse,
		OrderId:       warehouseOrder.OrderId,
		ProductStatus: warehouseOrder.ProductStatus,
		ProductMark:   warehouseOrder.ProductMark,
		CreatedAt:     warehouseOrder.CreatedAt,
		UpdatedAt:     warehouseOrder.UpdatedAt,
	}
	if err != nil {
		return warehouseOrdersResponse, err
	}
	return warehouseOrdersResponse, nil
}

// // Find all warehouse order data from database
func (service WarehouseOrderService) FindAllWarehouseOrder() ([]entity.WarehouseOrdersResponse, error) {
	warehouseOrders, err := service.warehouseOrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var warehouseOrdersResponse []entity.WarehouseOrdersResponse
	for _, warehouseOrder := range warehouseOrders {
		warehouseOrdersResponse = append(warehouseOrdersResponse, entity.WarehouseOrdersResponse{
			ID:            warehouseOrder.ID,
			WorkerUUID:    warehouseOrder.WorkerUUID,
			WarehouseId:   warehouseOrder.WarehouseId,
			Warehouse:     warehouseOrder.Warehouse,
			OrderId:       warehouseOrder.OrderId,
			ProductStatus: warehouseOrder.ProductStatus,
			ProductMark:   warehouseOrder.ProductMark,
			CreatedAt:     warehouseOrder.CreatedAt,
			UpdatedAt:     warehouseOrder.UpdatedAt,
		})
	}
	return warehouseOrdersResponse, nil
}

// Find warehouse order data by id from database
func (service WarehouseOrderService) FindWarehouseOrderById(id int) (entity.WarehouseOrdersResponse, error) {
	warehouseOrder, err := service.warehouseOrderRepository.FindByID(id)
	if err != nil {
		return entity.WarehouseOrdersResponse{}, err
	}
	warehouseOrdersResponse := entity.WarehouseOrdersResponse{
		ID:            warehouseOrder.ID,
		WorkerUUID:    warehouseOrder.WorkerUUID,
		WarehouseId:   warehouseOrder.WarehouseId,
		Warehouse:     warehouseOrder.Warehouse,
		OrderId:       warehouseOrder.OrderId,
		ProductStatus: warehouseOrder.ProductStatus,
		ProductMark:   warehouseOrder.ProductMark,
		CreatedAt:     warehouseOrder.CreatedAt,
		UpdatedAt:     warehouseOrder.UpdatedAt,
	}
	return warehouseOrdersResponse, nil

}

// Find warehouse order data by product status from database
func (service WarehouseOrderService) FindWarehouseOrderByProductStatus(productStatus string) ([]entity.WarehouseOrdersResponse, error) {
	warehouseOrder, err := service.warehouseOrderRepository.FindByProductStatus(productStatus)
	if err != nil {
		return nil, err
	}
	var warehouseOrdersResponse []entity.WarehouseOrdersResponse
	for _, warehouseOrder := range warehouseOrder {
		warehouseOrdersResponse = append(warehouseOrdersResponse, entity.WarehouseOrdersResponse{
			ID:            warehouseOrder.ID,
			WorkerUUID:    warehouseOrder.WorkerUUID,
			WarehouseId:   warehouseOrder.WarehouseId,
			Warehouse:     warehouseOrder.Warehouse,
			OrderId:       warehouseOrder.OrderId,
			ProductStatus: warehouseOrder.ProductStatus,
			ProductMark:   warehouseOrder.ProductMark,
			CreatedAt:     warehouseOrder.CreatedAt,
			UpdatedAt:     warehouseOrder.UpdatedAt,
		})
	}
	return warehouseOrdersResponse, nil
}

// Find warehouse order data by product mark from database
func (service WarehouseOrderService) FindWarehouseOrderByProductMark(productMark string) ([]entity.WarehouseOrdersResponse, error) {
	warehouseOrder, err := service.warehouseOrderRepository.FindByProductMark(productMark)
	if err != nil {
		return nil, err
	}
	var warehouseOrdersResponse []entity.WarehouseOrdersResponse
	for _, warehouseOrder := range warehouseOrder {
		warehouseOrdersResponse = append(warehouseOrdersResponse, entity.WarehouseOrdersResponse{
			ID:            warehouseOrder.ID,
			WorkerUUID:    warehouseOrder.WorkerUUID,
			WarehouseId:   warehouseOrder.WarehouseId,
			Warehouse:     warehouseOrder.Warehouse,
			OrderId:       warehouseOrder.OrderId,
			ProductStatus: warehouseOrder.ProductStatus,
			ProductMark:   warehouseOrder.ProductMark,
			CreatedAt:     warehouseOrder.CreatedAt,
			UpdatedAt:     warehouseOrder.UpdatedAt,
		})
	}
	return warehouseOrdersResponse, nil
}

// Update warehouse order data by id from database
func (service WarehouseOrderService) UpdateWarehouseOrderById(id int, warehouseReq entity.UpdateWarehouseOrders) (entity.WarehouseOrdersResponse, error) {
	wos := entity.WarehouseOrders{
		WorkerUUID:    warehouseReq.WorkerUUID,
		WarehouseId:   warehouseReq.WarehouseId,
		OrderId:       warehouseReq.OrderId,
		ProductStatus: warehouseReq.ProductStatus,
		ProductMark:   warehouseReq.ProductMark,
		UpdatedAt:     time.Now().Format(time.RFC3339Nano),
	}
	warehouseOrder, err := service.warehouseOrderRepository.UpdateByID(id, wos)
	warehouseOrdersResponse := entity.WarehouseOrdersResponse{
		ID:            warehouseOrder.ID,
		WorkerUUID:    warehouseOrder.WorkerUUID,
		WarehouseId:   warehouseOrder.WarehouseId,
		Warehouse:     warehouseOrder.Warehouse,
		OrderId:       warehouseOrder.OrderId,
		ProductStatus: warehouseOrder.ProductStatus,
		ProductMark:   warehouseOrder.ProductMark,
		CreatedAt:     warehouseOrder.CreatedAt,
		UpdatedAt:     warehouseOrder.UpdatedAt,
	}
	if err != nil {
		return warehouseOrdersResponse, err
	}
	return warehouseOrdersResponse, nil
}

// Delete warehouse order data by id from database
func (service WarehouseOrderService) DeleteWarehouseOrderById(id int) error {
	err := service.warehouseOrderRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
