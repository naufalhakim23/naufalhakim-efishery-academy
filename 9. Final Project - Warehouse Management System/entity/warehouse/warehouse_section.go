package entity

type WarehouseSection struct {
	ID               int       `gorm:"primaryKey;column:id" json:"id"`
	WarehouseID      int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse        Warehouse `gorm:"foreignKey:WarehouseID"`
	InventorySection int       `gorm:"column:inventory_section" json:"inventory_section"`
	InventoryAisle   int       `gorm:"column:inventory_aisle" json:"inventory_aisle"`
	InventoryRow     int       `gorm:"column:inventory_row" json:"inventory_row"`
	InventoryTier    int       `gorm:"column:inventory_tier" json:"inventory_tier"`
	CreatedAt        string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        string    `gorm:"column:updated_at" json:"updated_at"`
}

type CreateWarehouseSection struct {
	WarehouseID      int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse        Warehouse `gorm:"foreignKey:WarehouseID"`
	InventorySection int       `gorm:"column:inventory_section" json:"inventory_section"`
	InventoryAisle   int       `gorm:"column:inventory_aisle" json:"inventory_aisle"`
	InventoryRow     int       `gorm:"column:inventory_row" json:"inventory_row"`
	InventoryTier    int       `gorm:"column:inventory_tier" json:"inventory_tier"`
	CreatedAt        string    `gorm:"column:created_at" json:"created_at"`
}

type UpdateWarehouseSection struct {
	WarehouseID      int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse        Warehouse `gorm:"foreignKey:WarehouseID"`
	InventorySection int       `gorm:"column:inventory_section" json:"inventory_section"`
	InventoryAisle   int       `gorm:"column:inventory_aisle" json:"inventory_aisle"`
	InventoryRow     int       `gorm:"column:inventory_row" json:"inventory_row"`
	InventoryTier    int       `gorm:"column:inventory_tier" json:"inventory_tier"`
	UpdatedAt        string    `gorm:"column:updated_at" json:"updated_at"`
}

type WarehouseSectionResponse struct {
	ID               int       `gorm:"primaryKey;column:id" json:"id"`
	WarehouseID      int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse        Warehouse `gorm:"foreignKey:WarehouseID"`
	InventorySection int       `gorm:"column:inventory_section" json:"inventory_section"`
	InventoryAisle   int       `gorm:"column:inventory_aisle" json:"inventory_aisle"`
	InventoryRow     int       `gorm:"column:inventory_row" json:"inventory_row"`
	InventoryTier    int       `gorm:"column:inventory_tier" json:"inventory_tier"`
	CreatedAt        string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        string    `gorm:"column:updated_at" json:"updated_at"`
}
