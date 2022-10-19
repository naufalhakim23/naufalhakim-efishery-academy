package helpers

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id int, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(authHeader string) (interface{}, error) {
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	} else {
		return claims, nil
	}
}
