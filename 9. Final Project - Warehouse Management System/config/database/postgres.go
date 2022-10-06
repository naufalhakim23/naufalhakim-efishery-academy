package config

import (
	"fmt"
	"os"
	supplier "warehouse-management-system-eFishery/entity/supplier"
	warehouse "warehouse-management-system-eFishery/entity/warehouse"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("db_url")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
}

func Migrate() {
	DB.AutoMigrate(&warehouse.WarehouseWorkers{})
	DB.AutoMigrate(&warehouse.WarehouseRoles{})
	DB.AutoMigrate(&warehouse.WarehouseAuth{})
	DB.AutoMigrate(&warehouse.Warehouse{})
	DB.AutoMigrate(&warehouse.WarehouseAddress{})
	DB.AutoMigrate(&warehouse.WarehouseOrders{})
	DB.AutoMigrate(&warehouse.WarehouseProducts{})
	DB.AutoMigrate(&warehouse.WarehouseCategories{})

	DB.AutoMigrate(&supplier.Supplier{})
	DB.AutoMigrate(&supplier.SupplierAddress{})

}
