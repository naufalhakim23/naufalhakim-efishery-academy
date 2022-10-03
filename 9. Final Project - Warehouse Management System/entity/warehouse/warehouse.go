package entity

import "time"

type Warehouse struct {
	ID            int       `gorm:"primaryKey;column:id" json:"id"`
	WarehouseName string    `gorm:"column:warehouse_name" json:"warehouse_name"`
	WarehouseDesc string    `gorm:"column:warehouse_desc" json:"warehouse_desc"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type CreateWarehouse struct {
	WarehouseName string `gorm:"column:warehouse_name" json:"warehouse_name"`
	WarehouseDesc string `gorm:"column:warehouse_desc" json:"warehouse_desc"`
}

type UpdateWarehouse struct {
	WarehouseName string `gorm:"column:warehouse_name" json:"warehouse_name"`
	WarehouseDesc string `gorm:"column:warehouse_desc" json:"warehouse_desc"`
}

type WarehouseResponse struct {
	WarehouseName string `gorm:"column:warehouse_name" json:"warehouse_name"`
	WarehouseDesc string `gorm:"column:warehouse_desc" json:"warehouse_desc"`
}
