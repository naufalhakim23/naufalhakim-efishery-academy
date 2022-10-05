package repository

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseCategoriesRepository interface {
	Store(warehouseCategories entity.WarehouseCategories) (entity.WarehouseCategories, error)
	FindAll() ([]entity.WarehouseCategories, error)
	FindByID(id int) (entity.WarehouseCategories, error)
	Update(warehouseCategories entity.WarehouseCategories) (entity.WarehouseCategories, error)
	Delete(id int) error
}

type WarehouseCategoriesRepository struct {
	db *gorm.DB
}

func NewWarehouseCategoriesRepository(db *gorm.DB) *WarehouseCategoriesRepository {
	return &WarehouseCategoriesRepository{db}
}

// Storing warehouse categories data to database
func (wcr WarehouseCategoriesRepository) Store(warehouseCategories entity.WarehouseCategories) (entity.WarehouseCategories, error) {
	if err := wcr.db.Create(&warehouseCategories).Error; err != nil {
		return warehouseCategories, err
	}
	return warehouseCategories, nil
}

// Find all warehouse categories data from database
func (wcr WarehouseCategoriesRepository) FindAll() ([]entity.WarehouseCategories, error) {
	var warehouseCategories []entity.WarehouseCategories
	if err := wcr.db.Find(&warehouseCategories).Error; err != nil {
		return warehouseCategories, err
	}
	return warehouseCategories, nil
}

// Find by ID warehouse categories data from database
func (wcr WarehouseCategoriesRepository) FindByID(id int) (entity.WarehouseCategories, error) {
	var warehouseCategories entity.WarehouseCategories
	if err := wcr.db.Where("id = ?", id).Find(&warehouseCategories).Error; err != nil {
		return warehouseCategories, err
	}
	return warehouseCategories, nil
}

// Update warehouse categories data by id from database
func (wcr WarehouseCategoriesRepository) Update(warehouseCategories entity.WarehouseCategories) (entity.WarehouseCategories, error) {
	if err := wcr.db.Save(&warehouseCategories).Error; err != nil {
		return entity.WarehouseCategories{}, err
	}
	return warehouseCategories, nil
}

// Delete warehouse categories data by id from database
func (wcr WarehouseCategoriesRepository) Delete(id int) error {
	if err := wcr.db.Where("id = ?", id).Delete(&entity.WarehouseCategories{}).Error; err != nil {
		return err
	}
	return nil
}
