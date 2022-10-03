package repository

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseRepository interface {
	Store(warehouse entity.Warehouse) (entity.Warehouse, error)
	FindAll() ([]entity.Warehouse, error)
	FindByID(id int) (entity.Warehouse, error)
	Update(warehouse entity.Warehouse) (entity.Warehouse, error)
	Delete(id int) error
}

type WarehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepository {
	return &WarehouseRepository{db}
}

// Storing warehouse data to database
func (wr WarehouseRepository) Store(warehouse entity.Warehouse) (entity.Warehouse, error) {
	if err := wr.db.Create(&warehouse).Error; err != nil {
		return warehouse, err
	}
	return warehouse, nil
}

// Find all warehouse data from database
func (wr WarehouseRepository) FindAll() ([]entity.Warehouse, error) {
	var warehouses []entity.Warehouse
	if err := wr.db.Find(&warehouses).Error; err != nil {
		return warehouses, err
	}
	return warehouses, nil
}

// Find warehouse data by id from database
func (wr WarehouseRepository) FindByID(id int) (entity.Warehouse, error) {
	var warehouse entity.Warehouse
	if err := wr.db.Where("id = ?", id).Find(&warehouse).Error; err != nil {
		return warehouse, err
	}
	return warehouse, nil
}

// Update warehouse data by id from database
func (wr WarehouseRepository) Update(warehouse entity.Warehouse) (entity.Warehouse, error) {
	if err := wr.db.Save(&warehouse).Error; err != nil {
		return entity.Warehouse{}, err
	}
	return warehouse, nil
}

// Delete warehouse data by id from database
func (wr WarehouseRepository) Delete(id int) error {
	if err := wr.db.Where("id = ?", id).Delete(&entity.Warehouse{}).Error; err != nil {
		return err
	}
	return nil
}
