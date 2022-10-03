package entity

type WarehouseRoles struct {
	Id          int    `gorm:"primaryKey;column:id" json:"id"`
	Role        string `gorm:"column:role" json:"role"`
	Description string `gorm:"column:description" json:"description"`
}
type CreateWarehouseRoles struct {
	Role        string `gorm:"column:role" json:"role"`
	Description string `gorm:"column:description" json:"description"`
}
type UpdateWarehouseRoles struct {
	Role        string `gorm:"column:role" json:"role"`
	Description string `gorm:"column:description" json:"description"`
}
type WarehouseRolesResponse struct {
	Role        string `gorm:"column:role" json:"role"`
	Description string `gorm:"column:description" json:"description"`
}
