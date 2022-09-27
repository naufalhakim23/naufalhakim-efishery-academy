package config

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open("host=172.28.48.1 user=postgres password=postgres dbname=efishery-example port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
}

// func Migrate() {
// 	DB.AutoMigrate(&models.User{})
// }
