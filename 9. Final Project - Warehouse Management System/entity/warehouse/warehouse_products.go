package entity

type WarehouseProducts struct {
	ID             int    `gorm:"primaryKey;column:id" json:"id"`
	SKU            string `gorm:"column:sku" json:"sku"`
	ProductName    string `gorm:"column:product_name" json:"product_name"`
	ProductDesc    string `gorm:"column:product_desc" json:"product_desc"`
	Image          string `gorm:"column:image" json:"image"`
	Price          int    `gorm:"column:price" json:"price"`
	Stock          int    `gorm:"column:stock" json:"stock"`
	Weight         int    `gorm:"column:weight" json:"weight"`
	CategoryId     int    `gorm:"foreignkey;column:category_id" json:"category_id"`
	WarehouseId    int    `gorm:"foreignkey;column:warehouse_id" json:"warehouse_id"`
	SectionPlaceId int    `gorm:"foreignkey;column:section_place_id" json:"section_place_id"`
	SupplierId     int    `gorm:"foreignkey;column:supplier_id" json:"supplier_id"`
	CreatedAt      string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      string `gorm:"column:update_at" json:"update_at"`
}

type CreateWarehouseProducts struct {
	SKU            string `gorm:"column:sku" json:"sku"`
	ProductName    string `gorm:"column:product_name" json:"product_name"`
	ProductDesc    string `gorm:"column:product_desc" json:"product_desc"`
	Image          string `gorm:"column:image" json:"image"`
	Price          int    `gorm:"column:price" json:"price"`
	Stock          int    `gorm:"column:stock" json:"stock"`
	Weight         int    `gorm:"column:weight" json:"weight"`
	CategoryId     int    `gorm:"foreignkey;column:category_id" json:"category_id"`
	WarehouseId    int    `gorm:"foreignkey;column:warehouse_id" json:"warehouse_id"`
	SectionPlaceId int    `gorm:"foreignkey;column:section_place_id" json:"section_place_id"`
	SupplierId     int    `gorm:"foreignkey;column:supplier_id" json:"supplier_id"`
}

type UpdateWarehouseProducts struct {
	SKU            string `gorm:"column:sku" json:"sku"`
	ProductName    string `gorm:"column:product_name" json:"product_name"`
	ProductDesc    string `gorm:"column:product_desc" json:"product_desc"`
	Image          string `gorm:"column:image" json:"image"`
	Price          int    `gorm:"column:price" json:"price"`
	Stock          int    `gorm:"column:stock" json:"stock"`
	Weight         int    `gorm:"column:weight" json:"weight"`
	CategoryId     int    `gorm:"foreignkey;column:category_id" json:"category_id"`
	WarehouseId    int    `gorm:"foreignkey;column:warehouse_id" json:"warehouse_id"`
	SectionPlaceId int    `gorm:"foreignkey;column:section_place_id" json:"section_place_id"`
	SupplierId     int    `gorm:"foreignkey;column:supplier_id" json:"supplier_id"`
}

type WarehouseProductsResponse struct {
	SKU            string `gorm:"column:sku" json:"sku"`
	ProductName    string `gorm:"column:product_name" json:"product_name"`
	ProductDesc    string `gorm:"column:product_desc" json:"product_desc"`
	Image          string `gorm:"column:image" json:"image"`
	Price          int    `gorm:"column:price" json:"price"`
	Stock          int    `gorm:"column:stock" json:"stock"`
	Weight         int    `gorm:"column:weight" json:"weight"`
	CategoryId     int    `gorm:"foreignkey;column:category_id" json:"category_id"`
	WarehouseId    int    `gorm:"foreignkey;column:warehouse_id" json:"warehouse_id"`
	SectionPlaceId int    `gorm:"foreignkey;column:section_place_id" json:"section_place_id"`
	SupplierId     int    `gorm:"foreignkey;column:supplier_id" json:"supplier_id"`
}
