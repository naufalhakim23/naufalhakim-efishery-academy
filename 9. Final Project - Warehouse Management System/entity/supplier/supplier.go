package entity

type Supplier struct {
	ID           int    `gorm:"primaryKey;column:id" json:"id"`
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	CreateAt     string `gorm:"column:create_at" json:"create_at"`
	UpdateAt     string `gorm:"column:update_at" json:"update_at"`
}

type CreateSupplier struct {
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	CreateAt     string `gorm:"column:create_at" json:"create_at"`
}

type UpdateSupplier struct {
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	UpdateAt     string `gorm:"column:update_at" json:"update_at"`
}

type SupplierResponse struct {
	ID           int    `gorm:"primaryKey;column:id" json:"id"`
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	CreateAt     string `gorm:"column:create_at" json:"create_at"`
	UpdateAt     string `gorm:"column:update_at" json:"update_at"`
}
