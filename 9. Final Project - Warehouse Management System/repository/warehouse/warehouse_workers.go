package repository

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseWorkersRepository interface {
	Store(entity.WarehouseWorkers, entity.WarehouseAuth) (entity.WarehouseWorkers, error)
	FindAll() ([]entity.WarehouseWorkers, error)
	FindByUUID(uuid string) (entity.WarehouseWorkers, error)
	UpdateByUUID(uuid string, warehouseWorkers entity.WarehouseWorkers) (entity.WarehouseWorkers, error)
	DeleteByUUID(uuid string) error
}

type WarehouseWorkersRepository struct {
	db *gorm.DB
}

func NewWarehouseWorkersRepository(db *gorm.DB) *WarehouseWorkersRepository {
	return &WarehouseWorkersRepository{db}
}

// Storing warehouse workers data to database and to warehouse auth table
func (wwr WarehouseWorkersRepository) Store(warehouseWorkers entity.WarehouseWorkers, warehouseWorkerAuth entity.WarehouseAuth) (entity.WarehouseWorkers, error) {
	if err := wwr.db.Create(&warehouseWorkers).Error; err != nil {
		return warehouseWorkers, err
	}
	var warehouse []entity.Warehouse
	if err := wwr.db.Where("id = ?", warehouseWorkers.WarehouseId).Find(&warehouse).Error; err != nil {
		return warehouseWorkers, err
	}
	warehouseWorkers.Warehouse = warehouse[0]

	if err := wwr.db.Create(&warehouseWorkerAuth).Error; err != nil {
		return warehouseWorkers, err
	}
	return warehouseWorkers, nil

}

// Find all warehouse workers data
func (wwr WarehouseWorkersRepository) FindAll() ([]entity.WarehouseWorkers, error) {
	var warehouseWorkers []entity.WarehouseWorkers
	if err := wwr.db.Find(&warehouseWorkers).Error; err != nil {
		return warehouseWorkers, err
	}
	return warehouseWorkers, nil
}

// Find warehouse workers data by UUID
func (wwr WarehouseWorkersRepository) FindByUUID(uuid string) (entity.WarehouseWorkers, error) {
	var warehouseWorkers entity.WarehouseWorkers
	if err := wwr.db.Where("uuid = ?", uuid).Find(&warehouseWorkers).Error; err != nil {
		return entity.WarehouseWorkers{}, err
	}
	return warehouseWorkers, nil
}

// Update warehouse workers data by UUID
func (wwr WarehouseWorkersRepository) UpdateByUUID(uuid string, warehouseWorkers entity.WarehouseWorkers) (entity.WarehouseWorkers, error) {
	var warehouseWorker entity.WarehouseWorkers
	if err := wwr.db.Where("uuid = ?", uuid).Find(&warehouseWorker).Error; err != nil {
		return entity.WarehouseWorkers{}, err
	}
	if err := wwr.db.Model(&warehouseWorker).Updates(warehouseWorkers).Error; err != nil {
		return entity.WarehouseWorkers{}, err
	}
	return warehouseWorker, nil
}

// Delete warehouse workers data by UUID
func (wwr WarehouseWorkersRepository) DeleteByUUID(uuid string) error {
	var warehouseWorkerAuth entity.WarehouseAuth
	if err := wwr.db.Where("uuid = ?", uuid).Delete(&warehouseWorkerAuth).Error; err != nil {
		return err
	}
	var warehouseWorkers entity.WarehouseWorkers
	if err := wwr.db.Where("uuid = ?", uuid).Delete(&warehouseWorkers).Error; err != nil {
		return err
	}
	return nil
}
