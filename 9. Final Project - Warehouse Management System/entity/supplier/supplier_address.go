package entity

type SupplierAddress struct {
	ID          int    `gorm:"primaryKey;column:id" json:"id"`
	SupplierId  int    `gorm:"foreignKey;column:supplier_id" json:"supplier_id"`
	FullAddress string `gorm:"column:full_address" json:"full_address"`
	SubDistrict string `gorm:"column:sub_district" json:"sub_district"`
	District    string `gorm:"column:district" json:"district"`
	Province    string `gorm:"column:province" json:"province"`
	PostalCode  int    `gorm:"column:postal_code" json:"postal_code"`
}
