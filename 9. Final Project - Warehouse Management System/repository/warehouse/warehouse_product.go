package repository

import (
	supplier "warehouse-management-system-eFishery/entity/supplier"
	entity "warehouse-management-system-eFishery/entity/warehouse"

	"gorm.io/gorm"
)

type InterfaceWarehouseProductRepository interface {
	Store(entity.WarehouseProducts) (entity.WarehouseProducts, error)
	GetAll() ([]entity.WarehouseProducts, error)
	GetbyID(int) (entity.WarehouseProducts, error)
	GetByPrice(minPrice, maxPrice int) ([]entity.WarehouseProducts, error)
	Update(entity.WarehouseProducts) (entity.WarehouseProducts, error)
	Delete(id int) error
}

type WarehouseProductRepository struct {
	db *gorm.DB
}

func NewWarehouseProductRepository(db *gorm.DB) *WarehouseProductRepository {
	return &WarehouseProductRepository{db}
}

// Storing warehouse product data to database
func (wpr WarehouseProductRepository) Store(warehouseProduct entity.WarehouseProducts) (entity.WarehouseProducts, error) {
	var warehouseSection []entity.WarehouseSection
	if err := wpr.db.Where("id=?", warehouseProduct.SectionPlaceID).Find(&warehouseSection).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.WarehouseSection = warehouseSection[0]

	var warehouseCategory []entity.WarehouseCategories
	if err := wpr.db.Where("id=?", warehouseProduct.CategoryID).Find(&warehouseCategory).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.WarehouseCategories = warehouseCategory[0]

	var supplier []supplier.Supplier
	if err := wpr.db.Where("id=?", warehouseProduct.SupplierID).Find(&supplier).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.Supplier = supplier[0]

	var warehouse []entity.Warehouse
	if err := wpr.db.Where("id=?", warehouseProduct.WarehouseID).Find(&warehouse).Error; err != nil {
		return warehouseProduct, err
	}

	if err := wpr.db.Create(&warehouseProduct).Error; err != nil {
		return warehouseProduct, err
	}
	return warehouseProduct, nil
}

// Get all warehouse product data from database
func (wpr WarehouseProductRepository) GetAll() ([]entity.WarehouseProducts, error) {
	var warehouseProduct []entity.WarehouseProducts
	var warehouseSection []entity.WarehouseSection
	var warehouseCategory []entity.WarehouseCategories
	var supplier []supplier.Supplier
	var warehouse []entity.Warehouse
	if err := wpr.db.Find(&warehouseProduct).Error; err != nil {
		return warehouseProduct, err
	}
	for i := 0; i < len(warehouseProduct); i++ {
		if err := wpr.db.Where("id=?", warehouseProduct[i].SectionPlaceID).Find(&warehouseSection).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].WarehouseSection = warehouseSection[0]

		if err := wpr.db.Where("id=?", warehouseProduct[i].CategoryID).Find(&warehouseCategory).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].WarehouseCategories = warehouseCategory[0]

		if err := wpr.db.Where("id=?", warehouseProduct[i].SupplierID).Find(&supplier).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].Supplier = supplier[0]

		if err := wpr.db.Where("id=?", warehouseProduct[i].WarehouseID).Find(&warehouse).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].Warehouse = warehouse[0]
		// if err := wpr.db.Where("id = ?", warehouseSection[i].WarehouseID).Find(&warehouse).Error; err != nil {
		// 	return warehouseProduct, err
		// }
		// warehouseSection[i].Warehouse = warehouse[0]

	}
	return warehouseProduct, nil
}

// GetbyID warehouse product data from database
func (wpr WarehouseProductRepository) GetbyID(id int) (entity.WarehouseProducts, error) {
	var warehouseProduct entity.WarehouseProducts
	var warehouseSection []entity.WarehouseSection
	var warehouseCategory []entity.WarehouseCategories
	var supplier []supplier.Supplier
	var Warehouse []entity.Warehouse
	if err := wpr.db.Where("id=?", id).Find(&warehouseProduct).Error; err != nil {
		return warehouseProduct, err
	}
	if err := wpr.db.Where("id=?", warehouseProduct.SectionPlaceID).Find(&warehouseSection).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.WarehouseSection = warehouseSection[0]

	if err := wpr.db.Where("id=?", warehouseProduct.CategoryID).Find(&warehouseCategory).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.WarehouseCategories = warehouseCategory[0]

	if err := wpr.db.Where("id=?", warehouseProduct.SupplierID).Find(&supplier).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.Supplier = supplier[0]

	if err := wpr.db.Where("id=?", warehouseProduct.WarehouseID).Find(&Warehouse).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.Warehouse = Warehouse[0]
	return warehouseProduct, nil
}

// Get warehouse product data from database by price from minimum to maximum
func (wpr WarehouseProductRepository) GetByPrice(minPrice, maxPrice int) ([]entity.WarehouseProducts, error) {
	var warehouseProduct []entity.WarehouseProducts
	var warehouseSection []entity.WarehouseSection
	var warehouseCategory []entity.WarehouseCategories
	var supplier []supplier.Supplier
	var warehouse []entity.Warehouse
	if err := wpr.db.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&warehouseProduct).Error; err != nil {
		return warehouseProduct, err
	}
	for i := 0; i < len(warehouseProduct); i++ {
		if err := wpr.db.Where("id=?", warehouseProduct[i].SectionPlaceID).Find(&warehouseSection).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].WarehouseSection = warehouseSection[0]

		if err := wpr.db.Where("id=?", warehouseProduct[i].CategoryID).Find(&warehouseCategory).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].WarehouseCategories = warehouseCategory[0]

		if err := wpr.db.Where("id=?", warehouseProduct[i].SupplierID).Find(&supplier).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].Supplier = supplier[0]

		if err := wpr.db.Where("id=?", warehouseProduct[i].WarehouseID).Find(&warehouse).Error; err != nil {
			return warehouseProduct, err
		}
		warehouseProduct[i].Warehouse = warehouse[0]
	}
	return warehouseProduct, nil
}

// Update warehouse product data to database
func (wpr WarehouseProductRepository) Update(warehouseProduct entity.WarehouseProducts) (entity.WarehouseProducts, error) {
	var warehouseSection []entity.WarehouseSection
	if err := wpr.db.Where("id=?", warehouseProduct.SectionPlaceID).Find(&warehouseSection).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.WarehouseSection = warehouseSection[0]

	var warehouseCategory []entity.WarehouseCategories
	if err := wpr.db.Where("id=?", warehouseProduct.CategoryID).Find(&warehouseCategory).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.WarehouseCategories = warehouseCategory[0]

	var supplier []supplier.Supplier
	if err := wpr.db.Where("id=?", warehouseProduct.SupplierID).Find(&supplier).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.Supplier = supplier[0]

	var Warehouse []entity.Warehouse
	if err := wpr.db.Where("id=?", warehouseProduct.WarehouseID).Find(&Warehouse).Error; err != nil {
		return warehouseProduct, err
	}
	warehouseProduct.Warehouse = Warehouse[0]

	if err := wpr.db.Save(&warehouseProduct).Error; err != nil {
		return warehouseProduct, err
	}
	return warehouseProduct, nil
}

// Delete warehouse product data by id from database
func (wpr WarehouseProductRepository) Delete(id int) error {
	var warehouseProduct entity.WarehouseProducts
	if err := wpr.db.Where("id = ?", id).Delete(&warehouseProduct).Error; err != nil {
		return err
	}
	return nil
}
