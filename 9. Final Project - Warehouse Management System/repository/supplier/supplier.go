package repository

import (
	entity "warehouse-management-system-eFishery/entity/supplier"

	"gorm.io/gorm"
)

type InterfaceSupplierRepository interface {
	Store(supplier entity.Supplier) (entity.Supplier, error)
	FindAll() ([]entity.Supplier, error)
	FindByID(id int) (entity.Supplier, error)
	Update(supplier entity.Supplier) (entity.Supplier, error)
	Delete(id int) error
}

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{db}
}

// Storing supplier data to database
func (sr SupplierRepository) Store(supplier entity.Supplier) (entity.Supplier, error) {
	if err := sr.db.Create(&supplier).Error; err != nil {
		return supplier, err
	}
	return supplier, nil
}

// Find all supplier data from database
func (sr SupplierRepository) FindAll() ([]entity.Supplier, error) {
	var suppliers []entity.Supplier
	if err := sr.db.Find(&suppliers).Error; err != nil {
		return suppliers, err
	}
	return suppliers, nil
}

// Find by ID supplier data from database
func (sr SupplierRepository) FindByID(id int) (entity.Supplier, error) {
	var supplier entity.Supplier
	if err := sr.db.Where("id = ?", id).Find(&supplier).Error; err != nil {
		return supplier, err
	}
	return supplier, nil
}

// Update supplier data by id from database
func (sr SupplierRepository) Update(supplier entity.Supplier) (entity.Supplier, error) {
	if err := sr.db.Save(&supplier).Error; err != nil {
		return entity.Supplier{}, err
	}
	return supplier, nil
}

// Delete supplier data by id from database
func (sr SupplierRepository) Delete(id int) error {
	if err := sr.db.Where("id = ?", id).Delete(&entity.Supplier{}).Error; err != nil {
		return err
	}
	return nil
}
