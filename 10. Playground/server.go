package main

import (
	"fmt"
	"os"
	"playground/helpers"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println(secretKey)
	testJWT, _ := helpers.GenerateToken(1, "isozac@gmail.com")
	fmt.Println(testJWT)
	claims, _ := helpers.ValidateToken(testJWT)
	fmt.Println(claims)
}
