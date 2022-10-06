package entity

import (
	entity "warehouse-management-system-eFishery/entity/supplier"
)

type WarehouseProducts struct {
	ID                  int                 `gorm:"primaryKey;column:id" json:"id"`
	SKU                 string              `gorm:"column:sku" json:"sku"`
	ProductName         string              `gorm:"column:product_name" json:"product_name"`
	ProductDesc         string              `gorm:"column:product_desc" json:"product_desc"`
	Image               string              `gorm:"column:image" json:"image"`
	Price               int                 `gorm:"column:price" json:"price"`
	Stock               int                 `gorm:"column:stock" json:"stock"`
	Weight              int                 `gorm:"column:weight" json:"weight"`
	CategoryID          int                 `gorm:"column:category_id" json:"category_id"`
	WarehouseCategories WarehouseCategories `gorm:"foreignKey:CategoryID"`
	WarehouseID         int                 `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse           Warehouse           `gorm:"foreignKey:WarehouseID"`
	SectionPlaceID      int                 `gorm:"column:section_place_id" json:"section_place_id"`
	WarehouseSection    WarehouseSection    `gorm:"foreignKey:SectionPlaceID;references:ID"`
	SupplierID          int                 `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier            entity.Supplier     `gorm:"foreignKey:SupplierID" `
	CreatedAt           string              `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           string              `gorm:"column:updated_at" json:"updated_at"`
}

type CreateWarehouseProducts struct {
	SKU                 string              `gorm:"column:sku" json:"sku"`
	ProductName         string              `gorm:"column:product_name" json:"product_name"`
	ProductDesc         string              `gorm:"column:product_desc" json:"product_desc"`
	Image               string              `gorm:"column:image" json:"image"`
	Price               int                 `gorm:"column:price" json:"price"`
	Stock               int                 `gorm:"column:stock" json:"stock"`
	Weight              int                 `gorm:"column:weight" json:"weight"`
	CategoryID          int                 `gorm:"column:category_id" json:"category_id"`
	WarehouseCategories WarehouseCategories `gorm:"foreignKey:CategoryID"`
	WarehouseID         int                 `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse           Warehouse           `gorm:"foreignKey:WarehouseID"`
	SectionPlaceID      int                 `gorm:"column:section_place_id" json:"section_place_id"`
	WarehouseSection    WarehouseSection    `gorm:"foreignKey:SectionPlaceID;references:ID"`
	SupplierID          int                 `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier            entity.Supplier     `gorm:"foreignKey:SupplierID" `
	CreatedAt           string              `gorm:"column:created_at" json:"created_at"`
}

type UpdateWarehouseProducts struct {
	SKU                 string              `gorm:"column:sku" json:"sku"`
	ProductName         string              `gorm:"column:product_name" json:"product_name"`
	ProductDesc         string              `gorm:"column:product_desc" json:"product_desc"`
	Image               string              `gorm:"column:image" json:"image"`
	Price               int                 `gorm:"column:price" json:"price"`
	Stock               int                 `gorm:"column:stock" json:"stock"`
	Weight              int                 `gorm:"column:weight" json:"weight"`
	CategoryID          int                 `gorm:"column:category_id" json:"category_id"`
	WarehouseCategories WarehouseCategories `gorm:"foreignKey:CategoryID"`
	WarehouseID         int                 `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse           Warehouse           `gorm:"foreignKey:WarehouseID"`
	SectionPlaceID      int                 `gorm:"column:section_place_id" json:"section_place_id"`
	WarehouseSection    WarehouseSection    `gorm:"foreignKey:SectionPlaceID;references:ID"`
	SupplierID          int                 `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier            entity.Supplier     `gorm:"foreignKey:SupplierID" `
	UpdatedAt           string              `gorm:"column:updated_at" json:"updated_at"`
}

type WarehouseProductsResponse struct {
	SKU                 string              `gorm:"column:sku" json:"sku"`
	ProductName         string              `gorm:"column:product_name" json:"product_name"`
	ProductDesc         string              `gorm:"column:product_desc" json:"product_desc"`
	Image               string              `gorm:"column:image" json:"image"`
	Price               int                 `gorm:"column:price" json:"price"`
	Stock               int                 `gorm:"column:stock" json:"stock"`
	Weight              int                 `gorm:"column:weight" json:"weight"`
	CategoryID          int                 `gorm:"column:category_id" json:"category_id"`
	WarehouseCategories WarehouseCategories `gorm:"foreignKey:CategoryID"`
	WarehouseID         int                 `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse           Warehouse           `gorm:"foreignKey:WarehouseID"`
	SectionPlaceID      int                 `gorm:"column:section_place_id" json:"section_place_id"`
	WarehouseSection    WarehouseSection    `gorm:"foreignKey:SectionPlaceID;references:ID"`
	SupplierID          int                 `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier            entity.Supplier     `gorm:"foreignKey:SupplierID" `
	CreatedAt           string              `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           string              `gorm:"column:updated_at" json:"updated_at"`
}
