package repository

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseRolesRepository interface {
	Store(warehouseRoles entity.WarehouseRoles) (entity.WarehouseRoles, error)
	FindAll() ([]entity.WarehouseRoles, error)
	FindByID(id int) (entity.WarehouseRoles, error)
	Update(warehouseRoles entity.WarehouseRoles) (entity.WarehouseRoles, error)
	Delete(id int) error
}

type WarehouseRolesRepository struct {
	db *gorm.DB
}

func NewWarehouseRolesRepository(db *gorm.DB) *WarehouseRolesRepository {
	return &WarehouseRolesRepository{db}
}

// Storing warehouse roles data to database
func (wrr WarehouseRolesRepository) Store(warehouseRoles entity.WarehouseRoles) (entity.WarehouseRoles, error) {
	if err := wrr.db.Create(&warehouseRoles).Error; err != nil {
		return warehouseRoles, err
	}
	return warehouseRoles, nil
}

// Find all warehouse roles data from database
func (wrr WarehouseRolesRepository) FindAll() ([]entity.WarehouseRoles, error) {
	var warehouseRoles []entity.WarehouseRoles
	if err := wrr.db.Find(&warehouseRoles).Error; err != nil {
		return warehouseRoles, err
	}
	return warehouseRoles, nil
}

// Find warehouse roles data by id from database
func (wrr WarehouseRolesRepository) FindByID(id int) (entity.WarehouseRoles, error) {
	var warehouseRoles entity.WarehouseRoles
	if err := wrr.db.Where("id = ?", id).Find(&warehouseRoles).Error; err != nil {
		return warehouseRoles, err
	}
	return warehouseRoles, nil
}

// Update warehouse roles data by id from database
func (wrr WarehouseRolesRepository) Update(warehouseRoles entity.WarehouseRoles) (entity.WarehouseRoles, error) {
	if err := wrr.db.Save(&warehouseRoles).Error; err != nil {
		return entity.WarehouseRoles{}, err
	}
	return warehouseRoles, nil
}

// Delete warehouse roles data by id from database
func (wrr WarehouseRolesRepository) Delete(id int) error {
	if err := wrr.db.Where("id = ?", id).Delete(&entity.WarehouseRoles{}).Error; err != nil {
		return err
	}
	return nil
}
