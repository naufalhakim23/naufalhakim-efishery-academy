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

type (
	dsn struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
		SSLMode  string
		Timezone string
	}
)

func Connect() {
	dsn := dsn{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Timezone: os.Getenv("DB_TIMEZONE"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db_url := "host=" + dsn.Host + " user=" + dsn.User + " password=" + dsn.Password + " dbname=" + dsn.Dbname + " port=" + dsn.Port + " TimeZone=" + dsn.Timezone

	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})
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
