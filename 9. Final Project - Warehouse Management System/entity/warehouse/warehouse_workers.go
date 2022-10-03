package entity

type WarehouseWorkers struct {
	// ID             int            `gorm:"column:id" json:"id"`
	UUID           string         `gorm:"primaryKey;column:uuid" json:"uuid"`
	FirstName      string         `gorm:"column:first_name" json:"first_name"`
	LastName       string         `gorm:"column:last_name" json:"last_name"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	Email          string         `gorm:"column:email" json:"email"`
	WarehouseId    int            `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse      Warehouse      `gorm:"foreignKey:WarehouseId"`
	RolesId        int            `gorm:"column:roles_id" json:"roles_id"`
	WarehouseRoles WarehouseRoles `gorm:"foreignKey:RolesId"`
	CreatedAt      string         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      string         `gorm:"column:updated_at" json:"updated_at"`
}
type CreateWarehouseWorkers struct {
	// connect to worker auth gorm
	FirstName   string `gorm:"column:first_name" json:"first_name"`
	LastName    string `gorm:"column:last_name" json:"last_name"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Email       string `gorm:"column:email" json:"email"`
	Password    string `gorm:"column:password" json:"password"`
	WarehouseId int    `gorm:"column:warehouse_id;foreignKey" json:"warehouse_id"`
	RolesId     int    `gorm:"column:roles_id;foreignKey" json:"roles_id"`
}
type UpdateWarehouseWorkers struct {
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Email     string `gorm:"column:email" json:"email"`
}
type WarehouseWorkersResponse struct {
	UUID        string `gorm:"column:uuid" json:"uuid"`
	FirstName   string `gorm:"column:first_name" json:"first_name"`
	LastName    string `gorm:"column:last_name" json:"last_name"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Email       string `gorm:"column:email" json:"email"`
	WarehouseId int    `gorm:"column:warehouse_id;foreignKey" json:"warehouse_id"`
	RolesId     int    `gorm:"column:roles_id;foreignKey" json:"roles_id"`
}
