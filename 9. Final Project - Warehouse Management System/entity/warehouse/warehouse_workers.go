package entity

type WarehouseWorkers struct {
	UUID           string         `gorm:"primaryKey;column:uuid" json:"uuid"`
	FirstName      string         `gorm:"not null;column:first_name" json:"first_name"`
	LastName       string         `gorm:"not null;column:last_name" json:"last_name"`
	Phone          string         `gorm:"not null;column:phone;unique" json:"phone"`
	Email          string         `gorm:"not null;column:email;unique" json:"email"`
	Password       string         `gorm:"not null;column:password" json:"password"`
	WarehouseId    int            `gorm:"column:warehouse_id" json:"warehouse_id"`
	Warehouse      Warehouse      `gorm:"foreignKey:WarehouseId"`
	RolesId        int            `gorm:"column:roles_id" json:"roles_id"`
	WarehouseRoles WarehouseRoles `gorm:"foreignKey:RolesId"`
	CreatedAt      string         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      string         `gorm:"column:updated_at" json:"updated_at"`
}
type CreateWarehouseWorkers struct {
	// connect to worker auth gorm
	UUID        string `gorm:"primaryKey;column:uuid" json:"uuid"`
	FirstName   string `gorm:"column:first_name" json:"first_name"`
	LastName    string `gorm:"column:last_name" json:"last_name"`
	Phone       string `gorm:"not null;column:phone;unique" json:"phone"`
	Email       string `gorm:"column:email;unique" json:"email"`
	Password    string `gorm:"column:password" json:"password"`
	WarehouseId int    `gorm:"column:warehouse_id;foreignKey" json:"warehouse_id"`
	RolesId     int    `gorm:"column:roles_id;foreignKey" json:"roles_id"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
}
type UpdateWarehouseWorkers struct {
	FirstName string `gorm:"not null;column:first_name" json:"first_name"`
	LastName  string `gorm:"not null;column:last_name" json:"last_name"`
	Phone     string `gorm:"not null;column:phone;unique" json:"phone"`
	Email     string `gorm:"not null;column:email;unique" json:"email"`
	Password  string `gorm:"not null;column:password" json:"password"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}
type WarehouseWorkersResponse struct {
	UUID        string `gorm:"not null;column:uuid" json:"uuid"`
	FirstName   string `gorm:"not null;column:first_name" json:"first_name"`
	LastName    string `gorm:"not null;column:last_name" json:"last_name"`
	Phone       string `gorm:"not null;column:phone;unique" json:"phone"`
	Email       string `gorm:"not null;column:email;unique" json:"email"`
	WarehouseId int    `gorm:"not null;column:warehouse_id;foreignKey" json:"warehouse_id"`
	RolesId     int    `gorm:"not null;column:roles_id;foreignKey" json:"roles_id"`
}
