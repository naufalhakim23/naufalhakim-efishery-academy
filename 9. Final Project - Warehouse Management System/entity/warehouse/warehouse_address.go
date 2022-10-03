package entity

type WarehouseAddress struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	WarehouseID int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID"`
	FullAddress string    `gorm:"column:full_address" json:"full_address"`
	SubDistrict string    `gorm:"column:sub_district" json:"sub_district"`
	City        string    `gorm:"column:city" json:"city"`
	Province    string    `gorm:"column:province" json:"province"`
	PostalCode  int       `gorm:"column:postal_code" json:"postal_code"`
	Region      string    `gorm:"column:region" json:"region"`
	CreatedAt   string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   string    `gorm:"column:updated_at" json:"updated_at"`
}
type CreateWarehouseAddress struct {
	WarehouseID int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID"`
	FullAddress string    `gorm:"column:full_address" json:"full_address"`
	SubDistrict string    `gorm:"column:sub_district" json:"sub_district"`
	City        string    `gorm:"column:city" json:"city"`
	Province    string    `gorm:"column:province" json:"province"`
	Region      string    `gorm:"column:region" json:"region"`
	PostalCode  int       `gorm:"column:postal_code" json:"postal_code"`
	CreatedAt   string    `gorm:"column:created_at" json:"created_at"`
}
type UpdateWarehouseAddress struct {
	WarehouseID int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID"`
	FullAddress string    `gorm:"column:full_address" json:"full_address"`
	SubDistrict string    `gorm:"column:sub_district" json:"sub_district"`
	City        string    `gorm:"column:city" json:"city"`
	Province    string    `gorm:"column:province" json:"province"`
	PostalCode  int       `gorm:"column:postal_code" json:"postal_code"`
	Region      string    `gorm:"column:region" json:"region"`
	UpdatedAt   string    `gorm:"column:updated_at" json:"updated_at"`
}
type WarehouseAddressesResponse struct {
	ID          int       `gorm:"column:id" json:"id"`
	WarehouseID int       `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID"`
	FullAddress string    `gorm:"column:full_address" json:"full_address"`
	SubDistrict string    `gorm:"column:sub_district" json:"sub_district"`
	City        string    `gorm:"column:city" json:"city"`
	Province    string    `gorm:"column:province" json:"province"`
	PostalCode  int       `gorm:"column:postal_code" json:"postal_code"`
	Region      string    `gorm:"column:region" json:"region"`
	CreatedAt   string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   string    `gorm:"column:updated_at" json:"updated_at"`
}
