package repository

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseOrderRepository interface {
	Store(entity.WarehouseOrders) (entity.WarehouseOrders, error)
	FindAll() ([]entity.WarehouseOrders, error)
	FindByID(int) (entity.WarehouseOrders, error)
	FindByProductStatus(string) ([]entity.WarehouseOrders, error)
	FindByProductMark(string) ([]entity.WarehouseOrders, error)
	UpdateByID(int, entity.WarehouseOrders) (entity.WarehouseOrders, error)
	DeleteByID(id int) error
}

type WarehouseOrderRepository struct {
	db *gorm.DB
}

func NewWarehouseOrderRepository(db *gorm.DB) *WarehouseOrderRepository {
	return &WarehouseOrderRepository{db}
}

// Storing warehouse order data to database
func (wor WarehouseOrderRepository) Store(warehouseOrder entity.WarehouseOrders) (entity.WarehouseOrders, error) {
	var warehouse []entity.Warehouse
	if err := wor.db.Where("id = ?", warehouseOrder.WarehouseId).Find(&warehouse).Error; err != nil {
		return warehouseOrder, err
	}
	warehouseOrder.Warehouse = warehouse[0]
	if err := wor.db.Create(&warehouseOrder).Error; err != nil {
		return warehouseOrder, err
	}
	return warehouseOrder, nil
}

// Find all warehouse order data from database
func (wor WarehouseOrderRepository) FindAll() ([]entity.WarehouseOrders, error) {
	var warehouseOrders []entity.WarehouseOrders
	var warehouse []entity.Warehouse
	if err := wor.db.Find(&warehouseOrders).Error; err != nil {
		return warehouseOrders, err
	}

	// Find from id of warehouse from warehouse order
	for i := 0; i < len(warehouseOrders); i++ {
		if err := wor.db.Where("id = ?", warehouseOrders[i].WarehouseId).Find(&warehouse).Error; err != nil {
			return warehouseOrders, err
		}
		warehouseOrders[i].Warehouse = warehouse[0]
	}
	return warehouseOrders, nil
}

// Find warehouse order data by ID
func (wor WarehouseOrderRepository) FindByID(id int) (entity.WarehouseOrders, error) {
	var warehouseOrders entity.WarehouseOrders
	var warehouse []entity.Warehouse
	if err := wor.db.Where("id = ?", id).Find(&warehouseOrders).Error; err != nil {
		return warehouseOrders, err
	}
	if err := wor.db.Where("id = ?", warehouseOrders.WarehouseId).Find(&warehouse).Error; err != nil {
		return warehouseOrders, err
	}
	warehouseOrders.Warehouse = warehouse[0]
	return warehouseOrders, nil
}

// Find product status in order
func (wor WarehouseOrderRepository) FindByProductStatus(productStatus string) ([]entity.WarehouseOrders, error) {
	var warehouseOrders []entity.WarehouseOrders
	var warehouse []entity.Warehouse
	if err := wor.db.Where("product_status = ?", productStatus).Find(&warehouseOrders).Error; err != nil {
		return warehouseOrders, err
	}
	for i := 0; i < len(warehouseOrders); i++ {
		if err := wor.db.Where("id = ?", warehouseOrders[i].WarehouseId).Find(&warehouse).Error; err != nil {
			return warehouseOrders, err
		}
		warehouseOrders[i].Warehouse = warehouse[0]
	}
	return warehouseOrders, nil
}

// Find warehouse order data by product marking in or out
func (wor WarehouseOrderRepository) FindByProductMark(productMark string) ([]entity.WarehouseOrders, error) {
	var warehouseOrders []entity.WarehouseOrders
	var warehouse []entity.Warehouse
	if err := wor.db.Where("product_mark = ?", productMark).Find(&warehouseOrders).Error; err != nil {
		return warehouseOrders, err
	}
	for i := 0; i < len(warehouseOrders); i++ {
		if err := wor.db.Where("id = ?", warehouseOrders[i].WarehouseId).Find(&warehouse).Error; err != nil {
			return warehouseOrders, err
		}
		warehouseOrders[i].Warehouse = warehouse[0]
	}
	return warehouseOrders, nil
}

// Update warehouse order data by ID
func (wor WarehouseOrderRepository) UpdateByID(id int, warehouseOrder entity.WarehouseOrders) (entity.WarehouseOrders, error) {
	if err := wor.db.Model(&warehouseOrder).Where("id = ?", id).Updates(warehouseOrder).Error; err != nil {
		return warehouseOrder, err
	}
	var warehouse []entity.Warehouse
	if err := wor.db.Where("id = ?", warehouseOrder.WarehouseId).Find(&warehouse).Error; err != nil {
		return warehouseOrder, err
	}
	warehouseOrder.Warehouse = warehouse[0]

	return warehouseOrder, nil
}

// Delete warehouse order data by ID
func (wor WarehouseOrderRepository) DeleteByID(id int) error {
	if err := wor.db.Where("id = ?", id).Delete(&entity.WarehouseOrders{}).Error; err != nil {
		return err
	}
	return nil
}
