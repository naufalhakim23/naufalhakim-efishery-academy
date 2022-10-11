package entity

type WarehouseOrders struct {
	ID            int              `gorm:"primaryKey;column:id" json:"id"`
	WorkerUUID    string           `gorm:"column:worker_uuid"`
	Worker        WarehouseWorkers `gorm:"foreignKey:WorkerUUID" json:"worker"`
	WarehouseId   int              `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse     Warehouse        `gorm:"foreignKey:WarehouseId"`
	OrderId       int              `gorm:"column:order_id" json:"order_id"`
	ProductStatus string           `gorm:"column:product_status" json:"product_status"`
	ProductMark   string           `gorm:"column:product_mark" json:"product_mark"`
	CreatedAt     string           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     string           `gorm:"column:updated_at" json:"updated_at"`
}

type CreateWarehouseOrders struct {
	WorkerUUID    string `gorm:"column:worker_uuid" json:"worker_uuid"`
	WarehouseId   int    `gorm:"column:warehouse_id" json:"warehouse_id"`
	OrderId       int    `gorm:"column:order_id" json:"order_id"`
	ProductStatus string `gorm:"column:product_status" json:"product_status"`
	ProductMark   string `gorm:"column:product_mark" json:"product_mark"`
	CreatedAt     string `gorm:"column:created_at" json:"created_at"`
}

type UpdateWarehouseOrders struct {
	WorkerUUID    string `gorm:"column:worker_uuid" json:"worker_uuid"`
	WarehouseId   int    `gorm:"column:warehouse_id" json:"warehouse_id"`
	OrderId       int    `gorm:"column:order_id" json:"order_id"`
	ProductStatus string `gorm:"column:product_status" json:"product_status"`
	ProductMark   string `gorm:"column:product_mark" json:"product_mark"`
	UpdatedAt     string `gorm:"column:updated_at" json:"updated_at"`
}

type WarehouseOrdersResponse struct {
	ID            int       `gorm:"primaryKey;column:id" json:"id"`
	WorkerUUID    string    `gorm:"column:worker_uuid" json:"worker_id"`
	Warehouse     Warehouse `gorm:"foreignKey:WarehouseId"`
	OrderId       int       `gorm:"column:order_id" json:"order_id"`
	ProductStatus string    `gorm:"column:product_status" json:"product_status"`
	ProductMark   string    `gorm:"column:product_mark" json:"product_mark"`
	CreatedAt     string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     string    `gorm:"column:updated_at" json:"updated_at"`
}
