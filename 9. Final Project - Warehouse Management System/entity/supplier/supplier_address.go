package entity

type SupplierAddress struct {
	ID          int      `gorm:"primaryKey;column:id" json:"id"`
	SupplierID  int      `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier    Supplier `gorm:"foreignKey:SupplierID;references:ID" `
	FullAddress string   `gorm:"column:full_address" json:"full_address"`
	SubDistrict string   `gorm:"column:sub_district" json:"sub_district"`
	District    string   `gorm:"column:district" json:"district"`
	City        string   `gorm:"column:city" json:"city"`
	Province    string   `gorm:"column:province" json:"province"`
	PostalCode  int      `gorm:"column:postal_code" json:"postal_code"`
	CreateAt    string   `gorm:"column:create_at" json:"create_at"`
	UpdateAt    string   `gorm:"column:update_at" json:"update_at"`
}

type CreateSupplierAddress struct {
	SupplierID  int      `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier    Supplier `gorm:"foreignKey:SupplierID;references:ID" `
	FullAddress string   `gorm:"column:full_address" json:"full_address"`
	SubDistrict string   `gorm:"column:sub_district" json:"sub_district"`
	District    string   `gorm:"column:district" json:"district"`
	City        string   `gorm:"column:city" json:"city"`
	Province    string   `gorm:"column:province" json:"province"`
	PostalCode  int      `gorm:"column:postal_code" json:"postal_code"`
	CreateAt    string   `gorm:"column:create_at" json:"create_at"`
}
type UpdateSupplierAddress struct {
	SupplierID  int      `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier    Supplier `gorm:"foreignKey:SupplierID;references:ID" `
	FullAddress string   `gorm:"column:full_address" json:"full_address"`
	SubDistrict string   `gorm:"column:sub_district" json:"sub_district"`
	District    string   `gorm:"column:district" json:"district"`
	City        string   `gorm:"column:city" json:"city"`
	Province    string   `gorm:"column:province" json:"province"`
	PostalCode  int      `gorm:"column:postal_code" json:"postal_code"`
	UpdateAt    string   `gorm:"column:update_at" json:"update_at"`
}

type SupplierAddressResponse struct {
	ID          int      `gorm:"primaryKey;column:id" json:"id"`
	SupplierID  int      `gorm:"column:supplier_id" json:"supplier_id"`
	Supplier    Supplier `gorm:"foreignKey:SupplierID;references:ID" `
	FullAddress string   `gorm:"column:full_address" json:"full_address"`
	SubDistrict string   `gorm:"column:sub_district" json:"sub_district"`
	District    string   `gorm:"column:district" json:"district"`
	City        string   `gorm:"column:city" json:"city"`
	Province    string   `gorm:"column:province" json:"province"`
	PostalCode  int      `gorm:"column:postal_code" json:"postal_code"`
	CreateAt    string   `gorm:"column:create_at" json:"create_at"`
	UpdateAt    string   `gorm:"column:update_at" json:"update_at"`
}
