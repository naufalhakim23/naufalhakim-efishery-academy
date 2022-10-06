package entity

type Supplier struct {
	ID           int    `gorm:"primaryKey;column:id" json:"id"`
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	CreatedAt    string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    string `gorm:"column:updated_at" json:"updated_at"`
}

type CreateSupplier struct {
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	CreateAt     string `gorm:"column:created_at" json:"created_at"`
}

type UpdateSupplier struct {
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	UpdatedAt    string `gorm:"column:updated_at" json:"updated_at"`
}

type SupplierResponse struct {
	ID           int    `gorm:"primaryKey;column:id" json:"id"`
	SupplierName string `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierDesc string `gorm:"column:supplier_desc" json:"supplier_desc"`
	Phone        string `gorm:"column:phone" json:"phone"`
	Email        string `gorm:"column:email" json:"email"`
	CreatedAt    string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    string `gorm:"column:updated_at" json:"updated_at"`
}
