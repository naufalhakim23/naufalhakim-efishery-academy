package repository

import (
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseAddressRepository interface {
	Store(warehouseAddress entity.WarehouseAddress) (entity.WarehouseAddress, error)
	FindAll() ([]entity.WarehouseAddress, error)
	FindByID(id int) (entity.WarehouseAddress, error)
	Update(warehouseAddress entity.WarehouseAddress) (entity.WarehouseAddress, error)
	Delete(id int) error
}

type WarehouseAddressRepository struct {
	db *gorm.DB
}

func NewWarehouseAddressRepository(db *gorm.DB) *WarehouseAddressRepository {
	return &WarehouseAddressRepository{db}
}

// Creating warehouse address data to database
func (war WarehouseAddressRepository) Store(warehouseAddress entity.WarehouseAddress) (entity.WarehouseAddress, error) {
	var warehouse []entity.Warehouse
	if err := war.db.Where("id = ?", warehouseAddress.WarehouseID).Find(&warehouse).Error; err != nil {
		return warehouseAddress, err
	}
	warehouseAddress.Warehouse = warehouse[0]
	if err := war.db.Create(&warehouseAddress).Error; err != nil {
		return warehouseAddress, err
	}
	return warehouseAddress, nil
}

// Find all warehouse address data from database
func (war WarehouseAddressRepository) FindAll() ([]entity.WarehouseAddress, error) {
	var warehouseAddresses []entity.WarehouseAddress
	var warehouse []entity.Warehouse
	if err := war.db.Find(&warehouseAddresses).Error; err != nil {
		return warehouseAddresses, err
	}
	for i := 0; i < len(warehouseAddresses); i++ {
		if err := war.db.Where("id = ?", warehouseAddresses[i].WarehouseID).Find(&warehouse).Error; err != nil {
			return warehouseAddresses, err
		}
		warehouseAddresses[i].Warehouse = warehouse[0]
	}
	return warehouseAddresses, nil
}

// Find warehouse address data by id from database
func (war WarehouseAddressRepository) FindByID(id int) (entity.WarehouseAddress, error) {
	var warehouseAddress entity.WarehouseAddress
	var warehouse []entity.Warehouse
	if err := war.db.Where("id = ?", id).Find(&warehouseAddress).Error; err != nil {
		return warehouseAddress, err
	}
	if err := war.db.Where("id = ?", warehouseAddress.WarehouseID).Find(&warehouse).Error; err != nil {
		return warehouseAddress, err
	}
	warehouseAddress.Warehouse = warehouse[0]

	return warehouseAddress, nil
}

// Update warehouse address data by id from database
// Function Update still not working
func (war WarehouseAddressRepository) Update(warehouseAddress entity.WarehouseAddress) (entity.WarehouseAddress, error) {
	var warehouse []entity.Warehouse
	if err := war.db.Save(&warehouseAddress).Error; err != nil {
		return entity.WarehouseAddress{}, err
	}
	if err := war.db.Where("id = ?", warehouseAddress.WarehouseID).Find(&warehouse).Error; err != nil {
		return entity.WarehouseAddress{}, err
	}
	warehouseAddress.Warehouse = warehouse[0]
	return warehouseAddress, nil
}

// Delete warehouse address data by id from database
func (war WarehouseAddressRepository) Delete(id int) error {
	if err := war.db.Where("id = ?", id).Delete(&entity.WarehouseAddress{}).Error; err != nil {
		return err
	}
	return nil
}
