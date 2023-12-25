package service

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Not used currently in the API, but implemented to gain better understanding for future projects
func GenerateToken(appName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"app_name":   appName,
		"created_at": time.Now(),
	})

	fmt.Println(os.Getenv("TOKEN_SECRET"))
	tokenString, err := token.SignedString(os.Getenv("TOKEN_SECRET"))
	if err != nil {
		return "", err
	}

	fmt.Printf("the newly generated token: %v\n", tokenString)
	return tokenString, nil
}

func GenerateIdentifier() string {
	token := rand.Float64() * 100000000
	stringToken := fmt.Sprintf("%.0f", token)
	return stringToken
}
