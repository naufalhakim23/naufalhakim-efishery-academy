package repository

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseSectionRepository interface {
	Store(entity.WarehouseSection) (entity.WarehouseSection, error)
	FindAll() ([]entity.WarehouseSection, error)
	FindByID(int) (entity.WarehouseSection, error)
	Update(entity.WarehouseSection) (entity.WarehouseSection, error)
	Delete(int) error
}

type WarehouseSectionRepository struct {
	db *gorm.DB
}

func NewWarehouseSectionRepository(db *gorm.DB) *WarehouseSectionRepository {
	return &WarehouseSectionRepository{db}
}

// Storing warehouse section data to database
func (wsr WarehouseSectionRepository) Store(warehouseSection entity.WarehouseSection) (entity.WarehouseSection, error) {
	var warehouse []entity.Warehouse
	if err := wsr.db.Where("id = ?", warehouseSection.WarehouseID).Find(&warehouse).Error; err != nil {
		return warehouseSection, err
	}
	warehouseSection.Warehouse = warehouse[0]
	if err := wsr.db.Create(&warehouseSection).Error; err != nil {
		return warehouseSection, err
	}
	return warehouseSection, nil
}

// Find all warehouse section data from database
func (wsr WarehouseSectionRepository) FindAll() ([]entity.WarehouseSection, error) {
	var warehouseSections []entity.WarehouseSection
	var warehouse []entity.Warehouse
	if err := wsr.db.Find(&warehouseSections).Error; err != nil {
		return warehouseSections, err
	}
	for i := 0; i < len(warehouseSections); i++ {
		if err := wsr.db.Where("id = ?", warehouseSections[i].WarehouseID).Find(&warehouse).Error; err != nil {
			return warehouseSections, err
		}
		warehouseSections[i].Warehouse = warehouse[0]
	}
	return warehouseSections, nil
}

// Find warehouse section data by id from database
func (wsr WarehouseSectionRepository) FindByID(id int) (entity.WarehouseSection, error) {
	var warehouseSection entity.WarehouseSection
	var warehouse []entity.Warehouse
	if err := wsr.db.Where("id = ?", id).Find(&warehouseSection).Error; err != nil {
		return warehouseSection, err
	}
	if err := wsr.db.Where("id = ?", warehouseSection.WarehouseID).Find(&warehouse).Error; err != nil {
		return warehouseSection, err
	}
	warehouseSection.Warehouse = warehouse[0]
	return warehouseSection, nil
}

// Update warehouse section data by id from database
func (wsr WarehouseSectionRepository) Update(warehouseSection entity.WarehouseSection) (entity.WarehouseSection, error) {
	var warehouse []entity.Warehouse
	if err := wsr.db.Where("id = ?", warehouseSection.WarehouseID).Find(&warehouse).Error; err != nil {
		return warehouseSection, err
	}
	warehouseSection.Warehouse = warehouse[0]
	if err := wsr.db.Save(&warehouseSection).Error; err != nil {
		return warehouseSection, err
	}
	return warehouseSection, nil
}

// Delete warehouse section data by id from database
func (wsr WarehouseSectionRepository) Delete(id int) error {
	var warehouseSection entity.WarehouseSection
	if err := wsr.db.Where("id = ?", id).Delete(&warehouseSection).Error; err != nil {
		return err
	}
	return nil
}
