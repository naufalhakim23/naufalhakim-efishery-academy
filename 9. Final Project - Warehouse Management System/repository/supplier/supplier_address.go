package repository

import (
	entity "warehouse-management-system-eFishery/entity/supplier"

	"gorm.io/gorm"
)

type InterfaceSupplierAddressRepository interface {
	Store(supplierAddress entity.SupplierAddress) (entity.SupplierAddress, error)
	FindAll() ([]entity.SupplierAddress, error)
	FindByID(id int) (entity.SupplierAddress, error)
	Update(supplierAddress entity.SupplierAddress) (entity.SupplierAddress, error)
	Delete(id int) error
}

type SupplierAddressRepository struct {
	db *gorm.DB
}

func NewSupplierAddressRepository(db *gorm.DB) *SupplierAddressRepository {
	return &SupplierAddressRepository{db}
}

// Storing supplier address data to database
func (sar SupplierAddressRepository) Store(supplierAddress entity.SupplierAddress) (entity.SupplierAddress, error) {
	var supplier []entity.Supplier
	if err := sar.db.Where("id = ?", supplierAddress.SupplierID).Find(&supplier).Error; err != nil {
		return supplierAddress, err
	}
	supplierAddress.Supplier = supplier[0]
	if err := sar.db.Create(&supplierAddress).Error; err != nil {
		return supplierAddress, err
	}
	return supplierAddress, nil
}

// Find all supplier address data from database
func (sar SupplierAddressRepository) FindAll() ([]entity.SupplierAddress, error) {
	var supplierAddresses []entity.SupplierAddress
	var supplier []entity.Supplier
	if err := sar.db.Find(&supplierAddresses).Error; err != nil {
		return supplierAddresses, err
	}
	for i := 0; i < len(supplierAddresses); i++ {
		if err := sar.db.Where("id = ?", supplierAddresses[i].SupplierID).Find(&supplier).Error; err != nil {
			return supplierAddresses, err
		}
		supplierAddresses[i].Supplier = supplier[0]
	}
	return supplierAddresses, nil
}

// Find supplier address data by id from database
func (sar SupplierAddressRepository) FindByID(id int) (entity.SupplierAddress, error) {
	var supplierAddress entity.SupplierAddress
	var supplier []entity.Supplier
	if err := sar.db.Where("id = ?", id).Find(&supplierAddress).Error; err != nil {
		return supplierAddress, err
	}
	if err := sar.db.Where("id = ?", supplierAddress.SupplierID).Find(&supplier).Error; err != nil {
		return supplierAddress, err
	}
	supplierAddress.Supplier = supplier[0]
	return supplierAddress, nil
}

// Update supplier address data by id from database
func (sar SupplierAddressRepository) Update(supplierAddress entity.SupplierAddress) (entity.SupplierAddress, error) {
	var supplierAddressData entity.SupplierAddress
	var supplier []entity.Supplier
	if err := sar.db.Where("id = ?", supplierAddress.ID).Find(&supplierAddressData).Error; err != nil {
		return supplierAddress, err
	}
	if err := sar.db.Where("id = ?", supplierAddress.SupplierID).Find(&supplier).Error; err != nil {
		return supplierAddress, err
	}
	supplierAddress.Supplier = supplier[0]
	if err := sar.db.Save(&supplierAddress).Error; err != nil {
		return supplierAddress, err
	}
	return supplierAddress, nil
}

// Delete supplier address data by id from database
func (sar SupplierAddressRepository) Delete(id int) error {
	if err := sar.db.Delete(&entity.SupplierAddress{}, id).Error; err != nil {
		return err
	}
	return nil
}
