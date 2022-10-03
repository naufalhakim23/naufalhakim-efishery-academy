package entity

type WarehouseAddresses struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	WarehouseId int       `gorm:"foreignKey;column:warehouse_id" json:"warehouse_id"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId"`
	FullAddress string    `gorm:"column:full_address" json:"full_address"`
	SubDistrict string    `gorm:"column:sub_district" json:"sub_district"`
	District    string    `gorm:"column:district" json:"district"`
	Province    string    `gorm:"column:province" json:"province"`
	PostalCode  int       `gorm:"column:postal_code" json:"postal_code"`
	Region      string    `gorm:"column:region" json:"region"`
	CreatedAt   string    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   string    `gorm:"column:updated_at" json:"updated_at"`
}
type CreateWarehouseAddresses struct {
	WarehouseId int    `gorm:"foreignKey;column:warehouse_id" json:"warehouse_id"`
	FullAddress string `gorm:"column:full_address" json:"full_address"`
	SubDistrict string `gorm:"column:sub_district" json:"sub_district"`
	District    string `gorm:"column:district" json:"district"`
	Province    string `gorm:"column:province" json:"province"`
	Region      string `gorm:"column:region" json:"region"`
	PostalCode  int    `gorm:"column:postal_code" json:"postal_code"`
}
type UpdateWarehouseAddresses struct {
	FullAddress string `gorm:"column:full_address" json:"full_address"`
	SubDistrict string `gorm:"column:sub_district" json:"sub_district"`
	District    string `gorm:"column:district" json:"district"`
	Province    string `gorm:"column:province" json:"province"`
	PostalCode  int    `gorm:"column:postal_code" json:"postal_code"`
	Region      string `gorm:"column:region" json:"region"`
}
type WarehouseAddressesResponse struct {
	WarehouseId int    `gorm:"foreignKey;column:warehouse_id" json:"warehouse_id"`
	FullAddress string `gorm:"column:full_address" json:"full_address"`
	SubDistrict string `gorm:"column:sub_district" json:"sub_district"`
	District    string `gorm:"column:district" json:"district"`
	Province    string `gorm:"column:province" json:"province"`
	PostalCode  int    `gorm:"column:postal_code" json:"postal_code"`
	Region      string `gorm:"column:region" json:"region"`
}
