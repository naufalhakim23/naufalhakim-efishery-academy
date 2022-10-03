package entity

type WarehouseOrders struct {
	ID            int              `gorm:"primaryKey;column:id" json:"id"`
	WorkerId      int              `gorm:"column:worker_id"`
	Worker        WarehouseWorkers `gorm:"foreignKey:WorkerId" json:"worker"`
	WarehouseId   int              `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse     Warehouse        `gorm:"foreignKey:WarehouseId"`
	OrderId       int              `gorm:"foreignKey;column:order_id" json:"order_id"`
	ProductStatus string           `gorm:"column:product_status" json:"product_status"`
	ProductMark   string           `gorm:"column:product_mark" json:"product_mark"`
	CreatedAt     string           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     string           `gorm:"column:updated_at" json:"updated_at"`
}

type CreateWarehouseOrders struct {
	WarehouseId   int    `gorm:"foreignKey;column:warehouse_id" json:"warehouse_id"`
	OrderId       int    `gorm:"foreignKey;column:order_id" json:"order_id"`
	ProductStatus string `gorm:"column:product_status" json:"product_status"`
	ProductMark   string `gorm:"column:product_mark" json:"product_mark"`
}

type UpdateWarehouseOrders struct {
	WarehouseId   int    `gorm:"foreignKey;column:warehouse_id" json:"warehouse_id"`
	OrderId       int    `gorm:"foreignKey;column:order_id" json:"order_id"`
	ProductStatus string `gorm:"column:product_status" json:"product_status"`
	ProductMark   string `gorm:"column:product_mark" json:"product_mark"`
}

type WarehouseOrdersResponse struct {
	WorkerId      int    `gorm:"foreignKey;column:worker_id" json:"worker_id"`
	WarehouseId   int    `gorm:"foreignKey;column:warehouse_id" json:"warehouse_id"`
	OrderId       int    `gorm:"foreignKey;column:order_id" json:"order_id"`
	ProductStatus string `gorm:"column:product_status" json:"product_status"`
	ProductMark   string `gorm:"column:product_mark" json:"product_mark"`
}
