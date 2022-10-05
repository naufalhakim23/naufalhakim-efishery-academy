package entity

type WarehouseCategories struct {
	ID           int    `gorm:"primaryKey;column:id" json:"id"`
	CategoryName string `gorm:"column:category_name" json:"category_name"`
	CategoryDesc string `gorm:"column:category_desc" json:"category_desc"`
}
type CreateWarehouseCategories struct {
	CategoryName string `gorm:"column:category_name" json:"category_name"`
	CategoryDesc string `gorm:"column:category_desc" json:"category_desc"`
}
type UpdateWarehouseCategories struct {
	CategoryName string `gorm:"column:category_name" json:"category_name"`
	CategoryDesc string `gorm:"column:category_desc" json:"category_desc"`
}
type WarehouseCategoriesResponse struct {
	CategoryName string `gorm:"column:category_name" json:"category_name"`
	CategoryDesc string `gorm:"column:category_desc" json:"category_desc"`
}
