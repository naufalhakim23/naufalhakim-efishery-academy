package entity

import (
	"time"
)

type WarehouseWorkers struct {
	Id          int       `gorm:"column:id" json:"id"`
	UUID        string    `gorm:"column:uuid;primaryKey" json:"uuid"`
	FirstName   string    `gorm:"column:first_name" json:"first_name"`
	LastName    string    `gorm:"column:last_name" json:"last_name"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	Email       string    `gorm:"column:email" json:"email"`
	WarehouseId int       `gorm:"column:warehouse_id;foreignKey" json:"warehouse_id"`
	WorkerId    int       `gorm:"column:worker_id;foreignKey" json:"worker_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
type CreateWarehouseWorkers struct {
	// connect to worker auth gorm
	FirstName   string `gorm:"column:first_name" json:"first_name"`
	LastName    string `gorm:"column:last_name" json:"last_name"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Email       string `gorm:"column:email" json:"email"`
	Password    string `gorm:"column:password" json:"password"`
	WarehouseId int    `gorm:"column:warehouse_id;foreignKey" json:"warehouse_id"`
	WorkerId    int    `gorm:"column:worker_id;foreignKey" json:"worker_id"`
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
	WorkerId    int    `gorm:"column:worker_id;foreignKey" json:"worker_id"`
}
